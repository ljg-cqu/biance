package main

import (
	"context"
	"fmt"
	"github.com/dgraph-io/ristretto"
	"github.com/ljg-cqu/biance/logger"
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
	// Create logger
	logger.DevMode = true
	logger.UseConsoleEncoder = true
	myLogger := logger.Default()

	var pricesCh = make(chan Prices, 512)
	var wg = new(sync.WaitGroup)

	priceTracker := PriceTracker{
		Logger:    myLogger,
		Interval:  time.Second * 1,
		PricesCh:  pricesCh,
		WaitGroup: wg,
	}

	tOneMinute, _ := new(big.Float).SetString("0.03")
	tThreeMinutes, _ := new(big.Float).SetString("0.05")
	tFiveMinutes, _ := new(big.Float).SetString("0.10")
	tFifteenMinutes, _ := new(big.Float).SetString("0.15")
	tHalfHour, _ := new(big.Float).SetString("0.20")
	tOneHour, _ := new(big.Float).SetString("0.25")
	tTwoHours, _ := new(big.Float).SetString("0.30")
	tFourHours, _ := new(big.Float).SetString("0.35")
	tSixHours, _ := new(big.Float).SetString("0.40")
	tEightHours, _ := new(big.Float).SetString("0.45")
	tTwelvesHours, _ := new(big.Float).SetString("0.50")
	tEighteenHours, _ := new(big.Float).SetString("0.55")
	tOneDay, _ := new(big.Float).SetString("0.60")
	tThreeDays, _ := new(big.Float).SetString("0.65")
	tFiveDays, _ := new(big.Float).SetString("0.70")
	tTenDays, _ := new(big.Float).SetString("0.75")
	tTwentyDays, _ := new(big.Float).SetString("0.80")
	tThirtyDays, _ := new(big.Float).SetString("0.85")

	threholds := map[Period]Threshold{
		PeriodOneMinute:      {PeriodOneMinute, tOneMinute},
		PeriodThreeMinutes:   {PeriodThreeMinutes, tThreeMinutes},
		PeriodFiveMinutes:    {PeriodFiveMinutes, tFiveMinutes},
		PeriodFifteenMinutes: {PeriodFifteenMinutes, tFifteenMinutes},
		PeriodHalfHour:       {PeriodHalfHour, tHalfHour},
		PeriodOneHour:        {PeriodOneHour, tOneHour},
		PeriodTwoHours:       {PeriodTwoHours, tTwoHours},
		PeriodFourHours:      {PeriodFourHours, tFourHours},
		PeriodSixHours:       {PeriodSixHours, tSixHours},
		PeriodEightHours:     {PeriodEightHours, tEightHours},
		PeriodTwelveHours:    {PeriodTwelveHours, tTwelvesHours},
		PeriodEighteenHours:  {PeriodEighteenHours, tEighteenHours},
		PeriodOneDay:         {PeriodOneDay, tOneDay},
		PeriodThreeDays:      {PeriodThreeDays, tThreeDays},
		PeriodFiveDays:       {PeriodFiveDays, tFiveDays},
		PeriodTenDays:        {PeriodTenDays, tTenDays},
		PeriodTwentyDays:     {PeriodTwentyDays, tTwentyDays},
		PeriodThirtyDays:     {PeriodThirtyDays, tThirtyDays},
	}

	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})

	priceHandler := PriceHandler{
		Logger:             myLogger,
		PricesCh:           pricesCh,
		WaitGroup:          wg,
		Thresholds:         threholds,
		Cache:              cache,
		CheckPriceInterval: time.Second * 5,
		MiniReportInterval: time.Minute * 5,
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
