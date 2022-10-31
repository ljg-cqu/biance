package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/dgraph-io/ristretto"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/asset"
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
	Cache     *ristretto.Cache

	client         biance.Client
	userAssetURL   string
	symbolPriceURL string

	convertThreshold     *big.Float
	lossPercentThreshold *big.Float
	checkPNLInterval     time.Duration

	reportCh           chan string
	miniReportInterval time.Duration
}

func (m *PNLMonitor) Init() {
	m.client = &http.Client{}
	m.userAssetURL = biance.URLs[biance.URLUserAsset]
	m.symbolPriceURL = biance.URLs[biance.URLSymbolPrice]
	m.checkPNLInterval = time.Second * 15
	m.reportCh = make(chan string, 1024)
	m.miniReportInterval = time.Second * 180
}

func (m *PNLMonitor) Run(ctx context.Context) {
	tCheck := time.NewTicker(m.checkPNLInterval)
	defer tCheck.Stop()
	defer m.WP.Done()

	//go m.sendPNLReportWith126Mail(ctx)
	go m.sendPNLReportWithQQMail(ctx)

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

			printGain, printLoss := buildPNLStr(freePNLsFilter, "5", "0.05")
			reportGain, reportLoss := buildPNLStr(freePNLsFilter, "10", "0.1")

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

func (m *PNLMonitor) sendPNLReportWith126Mail(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case content := <-m.reportCh:
			email := mail.NewMSG()
			email.SetFrom("Zealy <ljg_cqu@126.com>").
				AddTo("ljg_cqu@126.com", "1025003548@qq.com").
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
			if err != nil {
				m.Logger.ErrorOnError(err, "Failed to report price")
				continue
			}
		}
	}
}

func (m *PNLMonitor) sendPNLReportWithQQMail(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case content := <-m.reportCh:
			email := mail.NewMSG()
			email.SetFrom("Zealy <1025003548@qq.com>").
				AddTo("ljg_cqu@126.com", "1025003548@qq.com").
				SetSubject("Biance Investment PNL Report")
			email.SetBody(mail.TextPlain, content)

			err := backoff.RetryFnExponential10Times(m.Logger, ctx, time.Second, time.Second*10, func() (bool, error) {
				emailCli, err := smtp.NewEmailClient(smtp.QQMail, &tls.Config{InsecureSkipVerify: true}, "1025003548@qq.com", "ncoajiivbenpbfbh")
				if err != nil {
					return true, errors.Wrapf(err, "failed to create email client.")
				}
				err = emailCli.Send(email)
				if err != nil {
					return true, errors.Wrapf(err, "failed to send email")
				}
				return false, nil
			})
			if err != nil {
				m.Logger.ErrorOnError(err, "Failed to report price")
				continue
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
