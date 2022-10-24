package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/dgraph-io/ristretto"
	"github.com/ljg-cqu/biance/backoff"
	"github.com/ljg-cqu/biance/logger"
	"github.com/ljg-cqu/core/smtp"
	"github.com/pkg/errors"
	mail "github.com/xhit/go-simple-mail/v2"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	PeriodOneMinute      Period = "oneMinute"
	PeriodThreeMinutes   Period = "ThreeMinutes"
	PeriodFiveMinutes    Period = "fiveMinutes"
	PeriodFifteenMinutes Period = "fifteenMinutes"
	PeriodHalfHour       Period = "halfHour"
	PeriodOneHour        Period = "oneHour"
	PeriodTwoHours       Period = "twoHours"
	PeriodFourHours      Period = "fourHours"
	PeriodSixHours       Period = "sixHours"
	PeriodEightHours     Period = "eightHours"
	PeriodTwelveHours    Period = "twelveHours"
	PeriodEighteenHours  Period = "eighteenHours"
	PeriodOneDay         Period = "oneDay"
	PeriodThreeDays      Period = "threeDays"
	PeriodFiveDays       Period = "fiveDays"
	PeriodTenDays        Period = "tenDays"
	PeriodTwentyDays     Period = "twentyDays"
	PeriodThirtyDays     Period = "thirtyDays"
)

type Period string

type Threshold struct {
	Period Period
	T      *big.Float
}

type PriceChange struct {
	LatestPrice      Price // todo: consider comparing with average prices
	PriceDiffPercent *big.Float
	Threshold
}

type PricesChange []PriceChange

func (p PricesChange) Len() int {
	return len(p)
}

func (p PricesChange) Less(i, j int) bool {
	if p[i].PriceDiffPercent.Cmp(p[j].PriceDiffPercent) == -1 {
		return true
	}
	return false
}

func (p PricesChange) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PricesChange) Sort() {
	sort.Sort(p)
}

func (p PricesChange) String() string {
	var str string
	for i, price := range p {
		var pref string
		zeroFloat, _ := new(big.Float).SetString("0.00")
		switch zeroFloat.Cmp(price.PriceDiffPercent) {
		case -1:
			pref = "-"
		case 0:
			pref = "0"
		case 1:
			pref = "+"
		}
		str += fmt.Sprintf("(%v):%v | %v:%v | %v:%v\n", pref, i, price.LatestPrice.Symbol, price.LatestPrice.Price, price.Period, price.PriceDiffPercent)
	}
	return str
}

// ---

type PriceHandler struct {
	Logger             logger.Logger
	PricesCh           chan Prices
	WaitGroup          *sync.WaitGroup
	Thresholds         map[Period]Threshold
	Cache              *ristretto.Cache
	CheckPriceInterval time.Duration
	MiniReportInterval time.Duration // avoid report too frequently

	pricesHistory []Prices // todo: consider persistence and recovery
}

func (p *PriceHandler) Run(ctx context.Context) {
	t := time.NewTicker(time.Second * p.CheckPriceInterval)
	defer p.WaitGroup.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case prices := <-p.PricesCh:
			pricesBUSDUSDT := p.filterPricesBUSDUSDT(prices)
			pricesBUSDUSDT.Sort()
			if len(p.pricesHistory) > 5184000 { // 30*24*60*60*2
				p.pricesHistory = p.pricesHistory[2592000:] // 30*24*60*60
			}
			p.pricesHistory = append(p.pricesHistory, pricesBUSDUSDT)
		case <-t.C:
			p.Logger.Debug("Checking price change...")
			p1, p1_ := p.doCheckDifferences(PeriodOneMinute)
			p2, p2_ := p.doCheckDifferences(PeriodThreeMinutes)
			p3, p3_ := p.doCheckDifferences(PeriodFiveMinutes)
			p4, p4_ := p.doCheckDifferences(PeriodFifteenMinutes)
			p5, p5_ := p.doCheckDifferences(PeriodHalfHour)
			p6, p6_ := p.doCheckDifferences(PeriodOneHour)
			p7, p7_ := p.doCheckDifferences(PeriodTwoHours)
			p8, p8_ := p.doCheckDifferences(PeriodFourHours)
			p9, p9_ := p.doCheckDifferences(PeriodSixHours)
			p10, p10_ := p.doCheckDifferences(PeriodEightHours)
			p11, p11_ := p.doCheckDifferences(PeriodTwelveHours)
			p12, p12_ := p.doCheckDifferences(PeriodEighteenHours)
			p13, p13_ := p.doCheckDifferences(PeriodOneDay)
			p14, p14_ := p.doCheckDifferences(PeriodThreeDays)
			p15, p15_ := p.doCheckDifferences(PeriodFiveDays)
			p16, p16_ := p.doCheckDifferences(PeriodTenDays)
			p17, p17_ := p.doCheckDifferences(PeriodTwentyDays)
			p18, p18_ := p.doCheckDifferences(PeriodThirtyDays)

			pricesChangeToPrint := p.buildPricesChangeStr(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16, p17, p18)
			pricesChangeToReport := p.buildPricesChangeStr(p1_, p2_, p3_, p4_, p5_, p6_, p7_, p8_, p9_, p10_, p11_, p12_, p13_, p14_, p15_, p16_, p17_, p18_)

			p.Logger.InfoOnTrue(pricesChangeToPrint != "", pricesChangeToPrint)
			if pricesChangeToReport != "" {
				p.sendPriceChangeReport(ctx, pricesChangeToReport)
			}
		}
	}
}

func (p *PriceHandler) buildPricesChangeStr(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16, p17, p18 string) string {
	if p1 == "" && p2 == "" && p3 == "" && p4 == "" && p5 == "" && p6 == "" && p7 == "" && p8 == "" &&
		p9 == "" && p10 == "" && p11 == "" && p12 == "" && p13 == "" && p14 == "" && p15 == "" && p16 == "" && p17 == "" && p18 == "" {
		return ""
	}
	var priceChangeStr string
	if p1 != "" {
		priceChangeStr = fmt.Sprintf("---PeriodOneMinute\n%v\n\n", p1)
	}
	if p2 != "" {
		priceChangeStr += fmt.Sprintf("------PeriodThreeMinutes\n%v\n\n", p2)
	}
	if p3 != "" {
		priceChangeStr += fmt.Sprintf("---------PeriodFiveMinutes\n%v\n\n", p3)
	}
	if p4 != "" {
		priceChangeStr += fmt.Sprintf("------------PeriodFifteenMinutes\n%v\n\n", p4)
	}
	if p5 != "" {
		priceChangeStr += fmt.Sprintf("---------------PeriodHalfHour\n%v\n\n", p5)
	}
	if p6 != "" {
		priceChangeStr += fmt.Sprintf("------------------PeriodOneHour\n%v\n\n", p6)
	}
	if p7 != "" {
		priceChangeStr += fmt.Sprintf("---------------------PeriodTwoHoursn%v\n\n", p7)
	}
	if p8 != "" {
		priceChangeStr += fmt.Sprintf("------------------------PeriodFourHours\n%v\n\n", p8)
	}
	if p9 != "" {
		priceChangeStr += fmt.Sprintf("---------------------------PeriodSixHours\n%v\n\n", p8)
	}
	if p10 != "" {
		priceChangeStr += fmt.Sprintf("------------------------------PeriodEightHours\n%v\n\n", p9)
	}
	if p11 != "" {
		priceChangeStr += fmt.Sprintf("---------------------------------PeriodTwelvesHours\n%v\n\n", p9)
	}
	if p12 != "" {
		priceChangeStr += fmt.Sprintf("------------------------------------PeriodEighteenHours\n%v\n\n", p9)
	}
	if p13 != "" {
		priceChangeStr += fmt.Sprintf("---------------------------------------PeriodOneDay\n%v\n\n", p10)
	}
	if p14 != "" {
		priceChangeStr += fmt.Sprintf("------------------------------------------PeriodThreeDays\n%v\n", p11)
	}
	if p15 != "" {
		priceChangeStr += fmt.Sprintf("---------------------------------------------PeriodFiveDays\n%v\n\n", p12)
	}
	if p16 != "" {
		priceChangeStr += fmt.Sprintf("------------------------------------------------PeriodTenDaysn%v\n\n", p13)
	}
	if p17 != "" {
		priceChangeStr += fmt.Sprintf("---------------------------------------------------PeriodTwentyDays\n%v\n\n", p14)
	}
	if p18 != "" {
		priceChangeStr += fmt.Sprintf("------------------------------------------------------PeriodThirtyDays\n%v\n\n", p15)
	}
	return priceChangeStr
}

func (p *PriceHandler) sendPriceChangeReport(ctx context.Context, report string) {
	email := mail.NewMSG()
	email.SetFrom("Zealy <ljg_cqu@126.com>").
		AddTo("ljg_cqu@126.com").
		SetSubject("Biance Market Price Change Report")
	email.SetBody(mail.TextPlain, report)

	err := backoff.RetryFnExponential100Times(p.Logger, ctx, time.Second, time.Second*10, func() (bool, error) {
		emailCli := smtp.NewEmailClient(smtp.NetEase126Mail, &tls.Config{InsecureSkipVerify: true}, "ljg_cqu@126.com", "XROTXFGWZUILANPB")
		err := emailCli.Send(email)
		if err != nil {
			return true, errors.Wrapf(err, "failed to send email")
		}
		return false, nil
	})
	p.Logger.ErrorOnError(err, "Failed to report price")
}

func (p *PriceHandler) filterPricesBUSDUSDT(prices Prices) Prices {
	var pricesBUSDUSDT Prices
	var pricesUSDT Prices
	var pricesBUSDMap = make(map[string]Price)
	for _, price := range prices {
		switch {
		case strings.HasSuffix(price.Symbol, string(TokenBUSD)):
			pricesBUSDUSDT = append(pricesBUSDUSDT, price)
			token := strings.TrimSuffix(price.Symbol, string(TokenBUSD))
			pricesBUSDMap[token] = price
		case strings.HasSuffix(price.Symbol, string(TokenUSDT)):
			pricesUSDT = append(pricesUSDT, price)
		}
	}

	for _, priceUSDT := range pricesUSDT {
		token := strings.TrimSuffix(priceUSDT.Symbol, string(TokenUSDT))
		_, ok := pricesBUSDMap[token]
		if !ok {
			pricesBUSDUSDT = append(pricesBUSDUSDT, priceUSDT)
		}
	}
	return pricesBUSDUSDT
}

func (p *PriceHandler) doCheckDifferences(period Period) (toPrint string, toReport string) {
	threshold, ok := p.Thresholds[period]
	if !ok {
		return "", ""
	}
	pricesChange := p.checkDifferences(threshold)
	if len(pricesChange) == 0 {
		return "", ""
	}

	//var pricesChangeToReport PricesChange
	//for _, priceChange := range pricesChange {
	//	symbol := priceChange.LatestPrice.Symbol
	//	_, ok := p.Cache.Get(symbol)
	//	if ok {
	//		continue
	//	}
	//	pricesChangeToReport = append(pricesChangeToReport, priceChange)
	//	p.Cache.SetWithTTL(priceChange.LatestPrice.Symbol, "", 1, p.MiniReportInterval)
	//	p.Cache.Wait()
	//}

	pricesChange.Sort()
	//pricesChangeToReport.Sort()
	return pricesChange.String(), pricesChange.String()
}

func (p *PriceHandler) checkDifferences(threshold Threshold) PricesChange {
	pricesLen := len(p.pricesHistory)
	if pricesLen == 0 {
		return nil
	}
	pricesNew := p.pricesHistory[pricesLen-1]

	var pricesChange PricesChange

	var pricesOld Prices
	switch threshold.Period {
	case PeriodOneMinute:
		if len(p.pricesHistory) > 60 {
			pricesOld = p.pricesHistory[pricesLen-60]
		}
	case PeriodThreeMinutes:
		if len(p.pricesHistory) > 180 {
			pricesOld = p.pricesHistory[pricesLen-180]
		}
	case PeriodFiveMinutes:
		if len(p.pricesHistory) > 300 {
			pricesOld = p.pricesHistory[pricesLen-300]
		}
	case PeriodFifteenMinutes:
		if len(p.pricesHistory) > 900 {
			pricesOld = p.pricesHistory[pricesLen-900]
		}
	case PeriodHalfHour:
		if len(p.pricesHistory) > 1800 {
			pricesOld = p.pricesHistory[pricesLen-1800]
		}
	case PeriodOneHour:
		if len(p.pricesHistory) > 3600 {
			pricesOld = p.pricesHistory[pricesLen-3600]
		}
	case PeriodTwoHours:
		if len(p.pricesHistory) > 7200 {
			pricesOld = p.pricesHistory[pricesLen-7200]
		}
	case PeriodFourHours:
		if len(p.pricesHistory) > 14400 {
			pricesOld = p.pricesHistory[pricesLen-14400]
		}
	case PeriodSixHours:
		if len(p.pricesHistory) > 21600 {
			pricesOld = p.pricesHistory[pricesLen-21600]
		}
	case PeriodEightHours:
		if len(p.pricesHistory) > 28800 {
			pricesOld = p.pricesHistory[pricesLen-28800]
		}
	case PeriodTwelveHours:
		if len(p.pricesHistory) > 43200 {
			pricesOld = p.pricesHistory[pricesLen-43200]
		}
	case PeriodEighteenHours:
		if len(p.pricesHistory) > 64800 {
			pricesOld = p.pricesHistory[pricesLen-64800]
		}
	case PeriodOneDay:
		if len(p.pricesHistory) > 86400 {
			pricesOld = p.pricesHistory[pricesLen-86400]
		}
	case PeriodThreeDays:
		if len(p.pricesHistory) > 259200 {
			pricesOld = p.pricesHistory[pricesLen-259200]
		}
	case PeriodFiveDays:
		if len(p.pricesHistory) > 604800 {
			pricesOld = p.pricesHistory[pricesLen-604800]
		}
	case PeriodTenDays:
		if len(p.pricesHistory) > 864000 {
			pricesOld = p.pricesHistory[pricesLen-864000]
		}
	case PeriodTwentyDays:
		if len(p.pricesHistory) > 1728000 {
			pricesOld = p.pricesHistory[pricesLen-1728000]
		}
	case PeriodThirtyDays:
		if len(p.pricesHistory) > 2592000 {
			pricesOld = p.pricesHistory[pricesLen-2592000]
		}
	}

	var pricesOldMap = map[string]Price{}
	for _, price := range pricesOld {
		pricesOldMap[price.Symbol] = price
	}

	for _, priceNew := range pricesNew {
		if priceOld, ok := pricesOldMap[priceNew.Symbol]; ok {
			priceDiff := new(big.Float).Sub(priceNew.PriceFloat, priceOld.PriceFloat)
			priceDiffPercent := new(big.Float).Quo(priceDiff, priceOld.PriceFloat)
			if new(big.Float).Abs(priceDiffPercent).Cmp(threshold.T) == -1 {
				continue
			}

			priceChange := PriceChange{priceNew, priceDiffPercent, threshold}
			pricesChange = append(pricesChange, priceChange)
		}
	}

	return pricesChange
}
