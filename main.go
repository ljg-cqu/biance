package main

import (
	"context"
	"github.com/dgraph-io/ristretto"
	"github.com/ljg-cqu/biance/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Create logger
	logger.DevMode = true
	logger.UseConsoleEncoder = true
	myLogger := logger.Default()

	//var pricesCh = make(chan Prices, 512)
	var wg = new(sync.WaitGroup)
	//
	//priceTracker := PriceTracker{
	//	Logger:   myLogger,
	//	Interval: time.Second * 1,
	//	PricesCh: pricesCh,
	//	WP:       wg,
	//}
	//
	//tOneMinute, _ := new(big.Float).SetString("0.03")
	//tThreeMinutes, _ := new(big.Float).SetString("0.03")
	//tFiveMinutes, _ := new(big.Float).SetString("0.05")
	//tFifteenMinutes, _ := new(big.Float).SetString("0.05")
	//tHalfHour, _ := new(big.Float).SetString("0.10")
	//tOneHour, _ := new(big.Float).SetString("0.10")
	//tTwoHours, _ := new(big.Float).SetString("0.15")
	//tFourHours, _ := new(big.Float).SetString("0.15")
	//tSixHours, _ := new(big.Float).SetString("0.20")
	//tEightHours, _ := new(big.Float).SetString("0.20")
	//tTwelvesHours, _ := new(big.Float).SetString("0.25")
	//tEighteenHours, _ := new(big.Float).SetString("0.25")
	//tOneDay, _ := new(big.Float).SetString("0.30")
	//tThreeDays, _ := new(big.Float).SetString("0.30")
	//tFiveDays, _ := new(big.Float).SetString("0.35")
	//tTenDays, _ := new(big.Float).SetString("0.35")
	//tTwentyDays, _ := new(big.Float).SetString("0.40")
	//tThirtyDays, _ := new(big.Float).SetString("0.40")
	//
	//threholds := map[Period]Threshold{
	//	PeriodOneMinute:      {PeriodOneMinute, tOneMinute},
	//	PeriodThreeMinutes:   {PeriodThreeMinutes, tThreeMinutes},
	//	PeriodFiveMinutes:    {PeriodFiveMinutes, tFiveMinutes},
	//	PeriodFifteenMinutes: {PeriodFifteenMinutes, tFifteenMinutes},
	//	PeriodHalfHour:       {PeriodHalfHour, tHalfHour},
	//	PeriodOneHour:        {PeriodOneHour, tOneHour},
	//	PeriodTwoHours:       {PeriodTwoHours, tTwoHours},
	//	PeriodFourHours:      {PeriodFourHours, tFourHours},
	//	PeriodSixHours:       {PeriodSixHours, tSixHours},
	//	PeriodEightHours:     {PeriodEightHours, tEightHours},
	//	PeriodTwelveHours:    {PeriodTwelveHours, tTwelvesHours},
	//	PeriodEighteenHours:  {PeriodEighteenHours, tEighteenHours},
	//	PeriodOneDay:         {PeriodOneDay, tOneDay},
	//	PeriodThreeDays:      {PeriodThreeDays, tThreeDays},
	//	PeriodFiveDays:       {PeriodFiveDays, tFiveDays},
	//	PeriodTenDays:        {PeriodTenDays, tTenDays},
	//	PeriodTwentyDays:     {PeriodTwentyDays, tTwentyDays},
	//	PeriodThirtyDays:     {PeriodThirtyDays, tThirtyDays},
	//}
	//
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	//
	//miniReportThreshold, _ := new(big.Float).SetString("0.05")
	//priceHandler := PriceHandler{
	//	Logger:              myLogger,
	//	PricesCh:            pricesCh,
	//	WP:                  wg,
	//	Thresholds:          threholds,
	//	Cache:               cache,
	//	CheckPriceInterval:  time.Minute * 5,
	//	MiniReportInterval:  time.Minute * 5,
	//	MiniReportThreshold: miniReportThreshold,
	//}

	pnlMonitor := PNLMonitor{
		Logger:    myLogger,
		ApiKey:    "",
		SecretKey: "",
		WP:        wg,
		Cache:     cache,
	}
	pnlMonitor.Init()

	shutdownCtx, shutdown := context.WithCancel(context.Background())

	// Handle graceful shutdown.
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		shutdown()
	}()

	//wg.Add(1)
	//go priceTracker.Run(shutdownCtx)
	//
	//wg.Add(1)
	//go priceHandler.Run(shutdownCtx)

	wg.Add(1)
	go pnlMonitor.Run(shutdownCtx)

	wg.Wait()
}
