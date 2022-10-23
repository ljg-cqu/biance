package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"
)

const (
	TokenBUSD Token = "BUSD"
	TokenUSDT Token = "USDT"
)

type Token string

type Price struct {
	Symbol     string
	Price      string
	PriceFloat *big.Float
	When       time.Time `json:"-"`
}

type Prices []Price

func (p Prices) Len() int {
	return len(p)
}

func (p Prices) Less(i, j int) bool {
	if p[i].PriceFloat.Cmp(p[j].PriceFloat) == -1 {
		return true
	}
	return false
}

func (p Prices) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Prices) Sort() {
	sort.Sort(p)
}

func (p Prices) String() string {
	var str string
	for i, price := range p {
		str += fmt.Sprintf("%v %v:%v\n", i, price.Symbol, price.Price)
	}
	str += fmt.Sprintf("------------------\n")
	return str
}

// ---

func main() {
	var pricesCh = make(chan Prices, 512)
	var wg = new(sync.WaitGroup)

	priceTracker := PriceTracker{
		Interval:  time.Second * 1,
		PricesCh:  pricesCh,
		WaitGroup: wg,
	}

	tOneMinute, _ := new(big.Float).SetString("0.05")
	tThreeMinute, _ := new(big.Float).SetString("0.1")
	tFiveMinute, _ := new(big.Float).SetString("0.15")
	tFifteenMinute, _ := new(big.Float).SetString("0.2")
	tHalfHourMinute, _ := new(big.Float).SetString("0.25")
	tOneHour, _ := new(big.Float).SetString("0.3")
	tTwoHour, _ := new(big.Float).SetString("0.35")
	tFourHour, _ := new(big.Float).SetString("0.4")
	tEightHour, _ := new(big.Float).SetString("0.45")
	tOneDay, _ := new(big.Float).SetString("0.5")
	tThreeDays, _ := new(big.Float).SetString("0.55")
	tFiveDays, _ := new(big.Float).SetString("0.6")
	tTenDays, _ := new(big.Float).SetString("0.65")
	tTwentyDays, _ := new(big.Float).SetString("0.7")
	tThirtyDays, _ := new(big.Float).SetString("0.75")

	threholds := map[Period]Threshold{
		PeriodOneMinute:      {PeriodOneMinute, tOneMinute},
		PeriodFiveMinutes:    {PeriodFiveMinutes, tThreeMinute},
		PeriodThreeMinutes:   {PeriodThreeMinutes, tFiveMinute},
		PeriodFifteenMinutes: {PeriodFifteenMinutes, tFifteenMinute},
		PeriodHalfHour:       {PeriodHalfHour, tHalfHourMinute},
		PeriodOneHour:        {PeriodOneHour, tOneHour},
		PeriodTwoHours:       {PeriodTwoHours, tTwoHour},
		PeriodFourHours:      {PeriodFourHours, tFourHour},
		PeriodEightHours:     {PeriodEightHours, tEightHour},
		PeriodOneDay:         {PeriodOneDay, tOneDay},
		PeriodThreeDays:      {PeriodThreeDays, tThreeDays},
		PeriodFiveDays:       {PeriodFiveDays, tFiveDays},
		PeriodTenDays:        {PeriodTenDays, tTenDays},
		PeriodTwentyDays:     {PeriodTwentyDays, tTwentyDays},
		PeriodThirtyDays:     {PeriodThirtyDays, tThirtyDays},
	}

	priceHandler := PriceHandler{
		PricesCh:   pricesCh,
		WaitGroup:  wg,
		Thresholds: threholds,
	}

	shutdownCtx, shutdown := context.WithCancel(context.Background())

	// Handle graceful shutdown.
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		shutdown()
	}()

	wg.Add(1)
	go priceTracker.Run(shutdownCtx)

	wg.Add(1)
	go priceHandler.Run(shutdownCtx)

	wg.Wait()
}
