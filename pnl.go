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
	FilterLevel0 FilterLevel = iota
	FilterLevel1
	FilterLevel2
	FilterLevel3
	FilterLevel4
	FilterLevel5
	FilterLevel6
)

var FilterMap = Filter{
	FilterLevel0: {FilterLevel0, "0.03", "0.03", time.Second * 7, time.Minute * 7},
	FilterLevel1: {FilterLevel1, "0.05", "0.05", time.Second * 6, time.Minute * 6},
	FilterLevel2: {FilterLevel2, "0.10", "0.10", time.Second * 5, time.Minute * 5},
	FilterLevel3: {FilterLevel3, "0.15", "0.15", time.Second * 4, time.Minute * 4},
	FilterLevel4: {FilterLevel4, "0.20", "0.20", time.Second * 3, time.Minute * 3},
	FilterLevel5: {FilterLevel5, "0.25", "0.25", time.Second * 2, time.Minute * 2},
	FilterLevel6: {FilterLevel6, "0.30", "0.30", time.Second * 1, time.Minute * 1},
}

type Filter map[FilterLevel]FilterGainValAndLoss

type FilterGainValAndLoss struct {
	FilterLevel       FilterLevel
	GainPercent       string
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
	Filter    FilterGainValAndLoss

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
				key := string(freePNL.Token) + m.Filter.ReportPNLInterval.String()
				_, ok2 := m.Cache.Get(key)
				if ok1 || ok2 {
					continue
				}

				freePNLsFilter = append(freePNLsFilter, freePNL)
			}

			gain, loss := buildReport(freePNLsFilter, m.Filter.FilterLevel, m.Filter.GainPercent, m.Filter.LossPercent)
			content := gain + loss

			if gain != "" || loss != "" {
				fmt.Println(content)
				m.reportCh <- content
				for _, freePNLFilter := range freePNLsFilter {
					key := string(freePNLFilter.Token) + m.Filter.ReportPNLInterval.String()
					m.Cache.SetWithTTL(key, "", 1, m.Filter.ReportPNLInterval)
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
			m.Logger.DebugOnError(err, "Failed to send email with QQ mailbox.")
			err = email.SendPNLReportWith126Mail(m.Logger, ctx, subject, content)
			m.Logger.DebugOnError(err, "Failed to send email with 126 mailbox")
		}
	}
}

func buildReport(freePNLs pnl.FreePNLs, filterLevel FilterLevel, filterGainPercent, filterLossPercent string) (string, string) {
	var filterGainPNLs []pnl.FreePNL
	var filterLossPNLs []pnl.FreePNL
	zeroGain, _ := new(big.Float).SetString("0")
	zeroLoss, _ := new(big.Float).SetString("0")

	filterGainPercentBig, _ := new(big.Float).SetString(filterGainPercent)
	filterLossPercentBig, _ := new(big.Float).SetString(filterLossPercent)
	var totalGain = zeroGain
	var totalLoss = zeroLoss

	oneHundred, _ := new(big.Float).SetString("100")

	for _, freePNL := range freePNLs {
		if freePNL.PNLPercent.Cmp(filterGainPercentBig) == 1 {
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
		gainInfoStr = fmt.Sprintf("\n+++LEVEL:%v\n+++TOKEN:%v PROFIT:%v\n+++TIME:%v\n",
			filterLevel, len(filterGainPNLs), totalGain, time.Now())
		for i, gainPNL := range filterGainPNLs {
			gainInfoStr += fmt.Sprintf("(%v) %v %v => %v @%v%%\n", i, gainPNL.Symbol, gainPNL.PNLAmountConvertable,
				gainPNL.PNLValue, new(big.Float).Mul(oneHundred, gainPNL.PNLPercent))
		}
	}

	if len(filterLossPNLs) > 0 {
		lossInfoStr = fmt.Sprintf("\n---LEVEL:%v\n---TOKEN:%v LOSS:%v\n---%v\n",
			filterLevel, len(filterLossPNLs), totalLoss, time.Now())
		for i, lossPNL := range filterLossPNLs {
			lossInfoStr += fmt.Sprintf("(%v) %v %v @%v%%\n", i, lossPNL.Symbol,
				lossPNL.PNLValue, new(big.Float).Mul(oneHundred, lossPNL.PNLPercent))
		}
	}

	return gainInfoStr, lossInfoStr
}
