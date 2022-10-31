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

type PNLMonitor struct {
	Logger    logger.Logger
	ApiKey    string
	SecretKey string
	WP        *sync.WaitGroup
	Cache     *ristretto.Cache

	client         biance.Client
	userAssetURL   string
	symbolPriceURL string

	checkPNLInterval   time.Duration
	miniReportInterval time.Duration

	printFilterGainValue   string
	printFilterLossPercent string

	emailFilterGainValue   string
	emailFilterLossPercent string

	reportCh chan string
}

func (m *PNLMonitor) Init() {
	m.client = &http.Client{}
	m.userAssetURL = biance.URLs[biance.URLUserAsset]
	m.symbolPriceURL = biance.URLs[biance.URLSymbolPrice]

	m.checkPNLInterval = time.Second * 15
	m.miniReportInterval = time.Second * 150

	m.printFilterGainValue = "5"
	m.printFilterLossPercent = "0.05"

	m.emailFilterGainValue = "8"
	m.emailFilterLossPercent = "0.08"

	m.reportCh = make(chan string, 1024)
}

func (m *PNLMonitor) Run(ctx context.Context) {
	tCheck := time.NewTicker(m.checkPNLInterval)
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

			printGain, printLoss := buildPNLStr(freePNLsFilter, m.printFilterGainValue, m.printFilterLossPercent)
			reportGain, reportLoss := buildPNLStr(freePNLsFilter, m.emailFilterGainValue, m.emailFilterLossPercent)

			if printGain != "" || printLoss != "" {
				fmt.Println(printGain + printLoss)
			}

			if reportGain != "" || reportLoss != "" {
				m.reportCh <- reportGain + reportLoss
				for _, freePNLFilter := range freePNLsFilter {
					m.Cache.SetWithTTL(string(freePNLFilter.Token), "", 1, m.miniReportInterval)
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

func buildPNLStr(freePNLs pnl.FreePNLs, filterGainVal, filterLossPercent string) (string, string) {
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
