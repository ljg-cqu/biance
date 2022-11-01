package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/dgraph-io/ristretto"
	"github.com/ljg-cqu/biance/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// TODO: auto night mode with cron

func main() {
	var nightFlag = flag.Bool("night", false, "enable night mode")
	flag.Parse()

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

	apiKey := ""
	secretKey := ""

	pnlMonitor1 := PNLMonitor{
		Logger:    myLogger,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		WP:        wg,
		Cache:     cache,
		Filter:    FilterMap[FilterLevel1],
	}
	pnlMonitor1.Init()

	pnlMonitor2 := PNLMonitor{
		Logger:    myLogger,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		WP:        wg,
		Cache:     cache,
		Filter:    FilterMap[FilterLevel2],
	}
	pnlMonitor2.Init()

	pnlMonitor3 := PNLMonitor{
		Logger:    myLogger,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		WP:        wg,
		Cache:     cache,
		Filter:    FilterMap[FilterLevel3],
	}
	pnlMonitor3.Init()

	pnlMonitor4 := PNLMonitor{
		Logger:    myLogger,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		WP:        wg,
		Cache:     cache,
		Filter:    FilterMap[FilterLevel4],
	}
	pnlMonitor4.Init()

	pnlMonitor5 := PNLMonitor{
		Logger:    myLogger,
		ApiKey:    apiKey,
		SecretKey: secretKey,
		WP:        wg,
		Cache:     cache,
		Filter:    FilterMap[FilterLevel5],
	}
	pnlMonitor5.Init()

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

	normalPNLMonitors := func() {
		wg.Add(1)
		go pnlMonitor2.Run(shutdownCtx)

		wg.Add(1)
		go pnlMonitor3.Run(shutdownCtx)

		wg.Add(1)
		go pnlMonitor4.Run(shutdownCtx)

		wg.Add(1)
		go pnlMonitor5.Run(shutdownCtx)
	}
	if *nightFlag {
		fmt.Println("Monitor PNL in night mode")
		normalPNLMonitors()
	} else {
		fmt.Println("Monitor PNL in day mode")
		wg.Add(1)
		go pnlMonitor1.Run(shutdownCtx)
		normalPNLMonitors()
	}

	wg.Wait()
}
