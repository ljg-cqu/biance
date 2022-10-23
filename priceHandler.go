package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/ljg-cqu/core/smtp"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
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
	PeriodEightHours     Period = "eightHours"
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
		str += fmt.Sprintf("(%v) %v:%v | %v:%v\n", i, price.LatestPrice.Symbol, price.LatestPrice.Price, price.Period, price.PriceDiffPercent)
	}
	return str
}

// ---

type PriceHandler struct {
	PricesCh   chan Prices
	WaitGroup  *sync.WaitGroup
	Thresholds map[Period]Threshold

	pricesHistory []Prices // todo: consider persistence and recovery
}

func (p *PriceHandler) Run(ctx context.Context) {
	t := time.NewTicker(time.Second * 15)
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
			p1 := p.doCheckDifferences(PeriodOneMinute)
			p2 := p.doCheckDifferences(PeriodThreeMinutes)
			p3 := p.doCheckDifferences(PeriodFiveMinutes)
			p4 := p.doCheckDifferences(PeriodFifteenMinutes)
			p5 := p.doCheckDifferences(PeriodHalfHour)
			p6 := p.doCheckDifferences(PeriodOneHour)
			p7 := p.doCheckDifferences(PeriodTwoHours)
			p8 := p.doCheckDifferences(PeriodFourHours)
			p9 := p.doCheckDifferences(PeriodEightHours)
			p10 := p.doCheckDifferences(PeriodOneDay)
			p11 := p.doCheckDifferences(PeriodThreeDays)
			p12 := p.doCheckDifferences(PeriodFiveDays)
			p13 := p.doCheckDifferences(PeriodTenDays)
			p14 := p.doCheckDifferences(PeriodTwentyDays)
			p15 := p.doCheckDifferences(PeriodThirtyDays)

			if p1 == "" && p2 == "" && p3 == "" && p4 == "" && p5 == "" && p6 == "" && p7 == "" && p8 == "" &&
				p9 == "" && p10 == "" && p11 == "" && p12 == "" && p13 == "" && p14 == "" && p15 == "" {
				continue
			}

			var priceChangeReport string
			if p1 != "" {
				priceChangeReport = fmt.Sprintf("---PeriodOneMinute\n%v\n\n", p1)
			}
			if p2 != "" {
				priceChangeReport += fmt.Sprintf("------PeriodThreeMinutes\n%v\n\n", p2)
			}
			if p3 != "" {
				priceChangeReport += fmt.Sprintf("---------PeriodFiveMinutes\n%v\n\n", p3)
			}
			if p4 != "" {
				priceChangeReport += fmt.Sprintf("------------PeriodFifteenMinutes\n%v\n\n", p4)
			}
			if p5 != "" {
				priceChangeReport += fmt.Sprintf("---------------PeriodHalfHour\n%v\n\n", p5)
			}
			if p6 != "" {
				priceChangeReport += fmt.Sprintf("------------------PeriodOneHour\n%v\n\n", p6)
			}
			if p7 != "" {
				priceChangeReport += fmt.Sprintf("---------------------PeriodTwoHoursn%v\n\n", p7)
			}
			if p8 != "" {
				priceChangeReport += fmt.Sprintf("------------------------PeriodFourHours\n%v\n\n", p8)
			}
			if p9 != "" {
				priceChangeReport += fmt.Sprintf("---------------------------PeriodEightHours\n%v\n\n", p9)
			}
			if p10 != "" {
				priceChangeReport += fmt.Sprintf("------------------------------PeriodOneDay\n%v\n\n", p10)
			}
			if p11 != "" {
				priceChangeReport += fmt.Sprintf("---------------------------------PeriodThreeDays\n%v\n", p11)
			}
			if p12 != "" {
				priceChangeReport += fmt.Sprintf("------------------------------------PeriodFiveDays\n%v\n\n", p12)
			}
			if p13 != "" {
				priceChangeReport += fmt.Sprintf("---------------------------------------PeriodTenDaysn%v\n\n", p13)
			}
			if p14 != "" {
				priceChangeReport += fmt.Sprintf("------------------------------------------PeriodTwentyDays\n%v\n\n", p14)
			}
			if p15 != "" {
				priceChangeReport += fmt.Sprintf("---------------------------------------------PeriodThirtyDays\n%v\n\n", p15)
			}

			fmt.Println(priceChangeReport)
			go p.sendPriceChangeReport(priceChangeReport)
		}
	}
}

func (p *PriceHandler) sendPriceChangeReport(report string) {
	emailCli := smtp.NewEmailClient(smtp.NetEase126Mail, &tls.Config{InsecureSkipVerify: true}, "ljg_cqu@126.com", "XROTXFGWZUILANPB")
	email := mail.NewMSG()
	email.SetFrom("Zealy <ljg_cqu@126.com>").
		AddTo("ljg_cqu@126.com").
		SetSubject("Biance Market Price Change Report")

	email.SetBody(mail.TextPlain, report)
	err := emailCli.Send(email)
	if err != nil {
		log.Printf("Failed to send price change report:%v\n", err)
	}
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

func (p *PriceHandler) doCheckDifferences(period Period) string {
	threshold, ok := p.Thresholds[period]
	if !ok {
		return ""
	}
	pricesChange := p.checkDifferences(threshold)
	if len(pricesChange) == 0 {
		return ""
	}

	pricesChange.Sort()
	return pricesChange.String()
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
	case PeriodEightHours:
		if len(p.pricesHistory) > 28800 {
			pricesOld = p.pricesHistory[pricesLen-14400]
		}
	case PeriodOneDay:
		if len(p.pricesHistory) > 86400 {
			pricesOld = p.pricesHistory[pricesLen-14400]
		}
	case PeriodThreeDays:
		if len(p.pricesHistory) > 259200 {
			pricesOld = p.pricesHistory[pricesLen-14400]
		}
	case PeriodFiveDays:
		if len(p.pricesHistory) > 604800 {
			pricesOld = p.pricesHistory[pricesLen-14400]
		}
	case PeriodTenDays:
		if len(p.pricesHistory) > 864000 {
			pricesOld = p.pricesHistory[pricesLen-14400]
		}
	case PeriodTwentyDays:
		if len(p.pricesHistory) > 1728000 {
			pricesOld = p.pricesHistory[pricesLen-14400]
		}
	case PeriodThirtyDays:
		if len(p.pricesHistory) > 2592000 {
			pricesOld = p.pricesHistory[pricesLen-14400]
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
