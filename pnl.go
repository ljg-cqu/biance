package main

import (
	"context"
	"fmt"
	"github.com/dgraph-io/ristretto"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/asset"
	"github.com/ljg-cqu/biance/biance/pnl"
	"github.com/ljg-cqu/biance/email"
	"github.com/ljg-cqu/biance/logger"
	"math/big"
	"net/http"
	"sync"
	"time"
)

const (
	FilterLevelLow FilterLevel = iota
	FilterLevelMid
	FilterLevelHigh
	FilterLevelSuper
)

var FilterMap = Filter{
	FilterLevelLow:   {"5", "0.05", time.Second * 60, time.Minute * 15},
	FilterLevelMid:   {"10", "0.10", time.Second * 30, time.Minute * 5},
	FilterLevelHigh:  {"20", "0.15", time.Second * 15, time.Minute * 3},
	FilterLevelSuper: {"40", "0.20", time.Second * 1, time.Minute * 1},
}

type Filter map[FilterLevel]FilterGainValAndLossPercent

type FilterGainValAndLossPercent struct {
	GainVal           string
	LossPercent       string
	CheckPNLInterval  time.Duration
	ReportPNLInterval time.Duration
}

type FilterLevel int

type PNLMonitor struct {
	Logger    logger.Logger
	ApiKey    string
	SecretKey string
	WP        *sync.WaitGroup
	Cache     *ristretto.Cache
	Filter    FilterGainValAndLossPercent

	client         biance.Client
	userAssetURL   string
	symbolPriceURL string

	reportCh chan string
}

func (m *PNLMonitor) Init() {
	m.client = &http.Client{}
	m.userAssetURL = biance.URLs[biance.URLUserAsset]
	m.symbolPriceURL = biance.URLs[biance.URLSymbolPrice]
	m.reportCh = make(chan string, 1024)
}

func (m *PNLMonitor) Run(ctx context.Context) {
	tCheck := time.NewTicker(m.Filter.CheckPNLInterval)
	defer tCheck.Stop()
	defer m.WP.Done()

	go m.sendPNLReport(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		case <-tCheck.C:
			freePNLs, err := pnl.CheckFreePNLWithUSDTOrBUSD(m.client, m.userAssetURL, m.symbolPriceURL, "", m.ApiKey, m.SecretKey)
			if err != nil {
				m.Logger.ErrorOnError(err, "Failed to check PNL")
				continue
			}

			var freePNLsFilter pnl.FreePNLs

			var tokenFilerMap = map[asset.Token]string{
				"NBS":  "",
				"USDT": "",
				"BUSD": "",
				"VIDT": "",
				"OG":   "",
			}

			for _, freePNL := range freePNLs {
				_, ok1 := tokenFilerMap[freePNL.Token]
				_, ok2 := m.Cache.Get(string(freePNL.Token))
				if ok1 || ok2 {
					continue
				}

				freePNLsFilter = append(freePNLsFilter, freePNL)
			}

			gain, loss := buildReport(freePNLsFilter, m.Filter.GainVal, m.Filter.LossPercent)
			content := gain + loss

			if gain != "" || loss != "" {
				fmt.Println(content)
				m.reportCh <- content
				for _, freePNLFilter := range freePNLsFilter {
					m.Cache.SetWithTTL(string(freePNLFilter.Token), "", 1, m.Filter.ReportPNLInterval)
					m.Cache.Wait()
				}
			}
		}
	}
}

func (m *PNLMonitor) sendPNLReport(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case content := <-m.reportCh:
			subject := "Biance Investment PNL Report"
			err := email.SendPNLReportWithQQMail(m.Logger, ctx, subject, content)
			if err != nil {
				m.Logger.DebugOnError(err, "Failed to send email with QQ mailbox.")
				if err != nil {
					err := email.SendPNLReportWith126Mail(m.Logger, ctx, subject, content)
					m.Logger.ErrorOnError(err, "Failed to send email with 126 mailbox")
				}
			}
		}
	}
}

func buildReport(freePNLs pnl.FreePNLs, filterGainVal, filterLossPercent string) (string, string) {
	var filterGainPNLs []pnl.FreePNL
	var filterLossPNLs []pnl.FreePNL
	zeroGain, _ := new(big.Float).SetString("0")
	zeroLoss, _ := new(big.Float).SetString("0")

	filterGainValBig, _ := new(big.Float).SetString(filterGainVal)
	filterLossPercentBig, _ := new(big.Float).SetString(filterLossPercent)
	var totalGain = zeroGain
	var totalLoss = zeroLoss

	oneHundred, _ := new(big.Float).SetString("100")

	for _, freePNL := range freePNLs {
		if freePNL.PNLValue.Cmp(filterGainValBig) == 1 {
			filterGainPNLs = append(filterGainPNLs, freePNL)
			totalGain = new(big.Float).Add(totalGain, freePNL.PNLValue)
			continue
		}
		if new(big.Float).Abs(freePNL.PNLPercent).Cmp(filterLossPercentBig) == 1 {
			filterLossPNLs = append(filterLossPNLs, freePNL)
			totalLoss = new(big.Float).Add(totalLoss, freePNL.PNLValue)
		}
	}

	var gainInfoStr string
	var lossInfoStr string

	if len(filterGainPNLs) > 0 {
		gainInfoStr = fmt.Sprintf("\n++++++++++++++++++++tokens:%v profit:%v++++++++++++++++++++\n",
			len(filterGainPNLs), totalGain)
		for _, gainPNL := range filterGainPNLs {
			gainInfoStr += fmt.Sprintf("%v %v => %v @%v%%\n", gainPNL.Symbol, gainPNL.PNLAmountConvertable,
				gainPNL.PNLValue, new(big.Float).Mul(oneHundred, gainPNL.PNLPercent))
		}
	}

	if len(filterLossPNLs) > 0 {
		lossInfoStr = fmt.Sprintf("--------------------tokens:%v loss:%v--------------------\n",
			len(filterLossPNLs), totalLoss)
		for _, lossPNL := range filterLossPNLs {
			lossInfoStr += fmt.Sprintf("%v %v @%v%%\n", lossPNL.Symbol,
				lossPNL.PNLValue, new(big.Float).Mul(oneHundred, lossPNL.PNLPercent))
		}
	}

	return gainInfoStr, lossInfoStr
}
