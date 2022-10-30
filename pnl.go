package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/pnl"
	"github.com/ljg-cqu/biance/logger"
	"github.com/ljg-cqu/biance/utils/backoff"
	"github.com/ljg-cqu/core/smtp"
	"github.com/pkg/errors"
	mail "github.com/xhit/go-simple-mail/v2"
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

	client         biance.Client
	userAssetURL   string
	symbolPriceURL string

	convertThreshold     *big.Float
	lossPercentThreshold *big.Float
	checkPNLInterval     time.Duration

	miniPrintInterval time.Duration

	reportCh           chan string
	miniReportInterval time.Duration

	lastPrintTime  time.Time
	lastReportTime time.Time
}

func (m *PNLMonitor) Init() {
	m.client = &http.Client{}
	m.userAssetURL = biance.URLs[biance.URLUserAsset]
	m.symbolPriceURL = biance.URLs[biance.URLSymbolPrice]
	m.miniPrintInterval = time.Second * 15
	m.checkPNLInterval = time.Second * 5
	m.reportCh = make(chan string, 1024)
	m.miniReportInterval = time.Minute * 5
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

			printGain, printLoss := buildPNLStr(freePNLs, "1", "0.03")
			reportGain, reportLoss := buildPNLStr(freePNLs, "10", "0.1")

			if m.lastPrintTime.Add(m.miniPrintInterval).After(time.Now()) {
				if printGain != "" || printLoss != "" {
					fmt.Println(printGain + printLoss)
				}
				m.lastPrintTime = time.Now()
			}

			if m.lastReportTime.Add(m.miniReportInterval).After(time.Now()) {
				if reportGain != "" || reportLoss != "" {
					m.reportCh <- reportGain + reportLoss
				}
				m.lastReportTime = time.Now()
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
			email := mail.NewMSG()
			email.SetFrom("Zealy <ljg_cqu@126.com>").
				AddTo("ljg_cqu@126.com").
				SetSubject("Biance Investment PNL Report")
			email.SetBody(mail.TextPlain, content)

			err := backoff.RetryFnExponential10Times(m.Logger, ctx, time.Second, time.Second*10, func() (bool, error) {
				emailCli, err := smtp.NewEmailClient(smtp.NetEase126Mail, &tls.Config{InsecureSkipVerify: true}, "ljg_cqu@126.com", "XROTXFGWZUILANPB")
				if err != nil {
					return true, errors.Wrapf(err, "failed to create email client.")
				}
				err = emailCli.Send(email)
				if err != nil {
					return true, errors.Wrapf(err, "failed to send email")
				}
				return false, nil
			})
			m.Logger.ErrorOnError(err, "Failed to report price")
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
