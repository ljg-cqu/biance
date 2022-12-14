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
	FilterLevel0: {FilterLevel0, "0.03", "0.03", time.Second * 7, time.Second * 105},
	FilterLevel1: {FilterLevel1, "0.05", "0.05", time.Second * 6, time.Second * 90},
	FilterLevel2: {FilterLevel2, "0.10", "0.10", time.Second * 5, time.Second * 75},
	FilterLevel3: {FilterLevel3, "0.15", "0.15", time.Second * 4, time.Second * 60},
	FilterLevel4: {FilterLevel4, "0.20", "0.20", time.Second * 3, time.Second * 45},
	FilterLevel5: {FilterLevel5, "0.25", "0.25", time.Second * 2, time.Second * 30},
	FilterLevel6: {FilterLevel6, "0.30", "0.30", time.Second * 1, time.Second * 15},
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

	ReportGain bool
	ReportLoss bool

	emailGainReportCH chan string
	emailLossReportCH chan string

	gailEmailAddress string
	lossEmailAddress string

	lossReportTimeAmplifier time.Duration
}

func (m *PNLMonitor) Init() {
	m.client = &http.Client{}
	m.userAssetURL = biance.URLs[biance.URLUserAsset]
	m.symbolPriceURL = biance.URLs[biance.URLSymbolPrice]
	m.emailGainReportCH = make(chan string, 1024)
	m.emailLossReportCH = make(chan string, 1024)
	m.gailEmailAddress = "ljg_cqu@163.com"
	m.lossEmailAddress = "1025003548@qq.com"
	m.lossReportTimeAmplifier = 5
}

func (m *PNLMonitor) Run(ctx context.Context) {
	tCheck := time.NewTicker(m.Filter.CheckPNLInterval)
	defer tCheck.Stop()
	defer m.WP.Done()

	go m.sendPNLReportWIthEmail(ctx)

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
			if gain != "" || loss != "" {
				fmt.Println(gain + loss)
			}

			if gain != "" && m.ReportGain {
				m.emailGainReportCH <- gain
				for _, freePNLFilter := range freePNLsFilter {
					key := string(freePNLFilter.Token) + m.Filter.ReportPNLInterval.String()
					m.Cache.SetWithTTL(key, "", 1, m.Filter.ReportPNLInterval)
					m.Cache.Wait()
				}
			}

			if loss != "" && m.ReportLoss {
				m.emailLossReportCH <- loss
				for _, freePNLFilter := range freePNLsFilter {
					key := string(freePNLFilter.Token) + m.Filter.ReportPNLInterval.String()
					m.Cache.SetWithTTL(key, "", 1, m.Filter.ReportPNLInterval*m.lossReportTimeAmplifier)
					m.Cache.Wait()
				}
			}
		}
	}
}

func (m *PNLMonitor) sendPNLReportWIthEmail(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case content := <-m.emailGainReportCH:
			subject := "Biance Investment Gain Report"
			m.doSendPNLReportWIthEmail(ctx, subject, content, m.gailEmailAddress)
		case content := <-m.emailLossReportCH:
			subject := "Biance Investment Loss Report"
			m.doSendPNLReportWIthEmail(ctx, subject, content, m.lossEmailAddress)
		}
	}
}

func (m *PNLMonitor) doSendPNLReportWIthEmail(ctx context.Context, subject, content, to string) {
	err := email.Send163Mail(m.Logger, ctx, subject, content, to)
	m.Logger.DebugOnError(err, "Failed to send email with 163 mailbox")

	if err != nil {
		err = email.Send126Mail(m.Logger, ctx, subject, content, to)
		m.Logger.DebugOnError(err, "Failed to send email with 126 mailbox")
	}

	if err != nil {
		err := email.SendQQMail(m.Logger, ctx, subject, content, to)
		m.Logger.DebugOnError(err, "Failed to send email with QQ mailbox.")
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
	zero, _ := new(big.Float).SetString("0")
	for _, freePNL := range freePNLs {
		switch freePNL.PNLPercent.Cmp(zero) {
		case 1:
			if freePNL.PNLPercent.Cmp(filterGainPercentBig) == 1 {
				filterGainPNLs = append(filterGainPNLs, freePNL)
				totalGain = new(big.Float).Add(totalGain, freePNL.PNLValue)
			}
		case 0:
		case -1:
			if new(big.Float).Abs(freePNL.PNLPercent).Cmp(filterLossPercentBig) == 1 {
				filterLossPNLs = append(filterLossPNLs, freePNL)
				totalLoss = new(big.Float).Add(totalLoss, freePNL.PNLValue)
			}
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
