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

	tOneMinute, _ := new(big.Float).SetString("0.03")
	tThreeMinute, _ := new(big.Float).SetString("0.05")
	tFiveMinute, _ := new(big.Float).SetString("0.10")
	tFifteenMinute, _ := new(big.Float).SetString("0.15")
	tHalfHourMinute, _ := new(big.Float).SetString("0.20")
	tOneHour, _ := new(big.Float).SetString("0.25")
	tTwoHour, _ := new(big.Float).SetString("0.30")
	tFourHour, _ := new(big.Float).SetString("0.35")
	tEightHour, _ := new(big.Float).SetString("0.4")
	tOneDay, _ := new(big.Float).SetString("0.45")
	tThreeDays, _ := new(big.Float).SetString("0.50")
	tFiveDays, _ := new(big.Float).SetString("0.55")
	tTenDays, _ := new(big.Float).SetString("0.60")
	tTwentyDays, _ := new(big.Float).SetString("0.65")
	tThirtyDays, _ := new(big.Float).SetString("0.70")

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
