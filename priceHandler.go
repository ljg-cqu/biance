package main

import (
	"context"
	"fmt"
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
		str += fmt.Sprintf("%v %v:%v %v:%v\n", i, price.LatestPrice.Symbol, price.LatestPrice.Price, price.Period, price.PriceDiffPercent)
	}
	str += fmt.Sprintf("------------------\n")
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
			p.doCheckDifferences(PeriodOneMinute)
			p.doCheckDifferences(PeriodThreeMinutes)
			p.doCheckDifferences(PeriodFiveMinutes)
			p.doCheckDifferences(PeriodFifteenMinutes)
			p.doCheckDifferences(PeriodHalfHour)
			p.doCheckDifferences(PeriodOneHour)
			p.doCheckDifferences(PeriodTwoHours)
			p.doCheckDifferences(PeriodFourHours)
			p.doCheckDifferences(PeriodEightHours)
			p.doCheckDifferences(PeriodOneDay)
			p.doCheckDifferences(PeriodThreeDays)
			p.doCheckDifferences(PeriodFiveDays)
			p.doCheckDifferences(PeriodTenDays)
			p.doCheckDifferences(PeriodTwentyDays)
			p.doCheckDifferences(PeriodThirtyDays)
		}
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

func (p *PriceHandler) doCheckDifferences(period Period) {
	threshold, ok := p.Thresholds[period]
	if !ok {
		return
	}
	pricesChange := p.checkDifferences(threshold)
	if len(pricesChange) == 0 {
		return
	}

	pricesChange.Sort()
	fmt.Print(pricesChange)
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
