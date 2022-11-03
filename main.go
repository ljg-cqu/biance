package main

import (
	"context"
	"flag"
	"github.com/dgraph-io/ristretto"
	"github.com/ljg-cqu/biance/logger"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// TODO: auto night mode with cron

func main() {
	var levels = flag.String("reportLevels", "0 1 2 3 4 5 6", "report levels")
	flag.Parse()

	levelsSplit := strings.Fields(*levels)
	var levelsMap = make(map[string]bool)
	for _, level := range levelsSplit {
		levelsMap[level] = true
	}

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

	shutdownCtx, shutdown := context.WithCancel(context.Background())

	if levelsMap["0"] {
		pnlMonitor0 := PNLMonitor{
			Logger:    myLogger,
			ApiKey:    apiKey,
			SecretKey: secretKey,
			WP:        wg,
			Cache:     cache,
			Filter:    FilterMap[FilterLevel0],
		}
		pnlMonitor0.Init()
		wg.Add(1)
		go pnlMonitor0.Run(shutdownCtx)
		myLogger.Debug("Enable Level0 report")
	}

	if levelsMap["1"] {
		pnlMonitor1 := PNLMonitor{
			Logger:    myLogger,
			ApiKey:    apiKey,
			SecretKey: secretKey,
			WP:        wg,
			Cache:     cache,
			Filter:    FilterMap[FilterLevel1],
		}
		pnlMonitor1.Init()
		wg.Add(1)
		go pnlMonitor1.Run(shutdownCtx)
		myLogger.Debug("Enable Level1 report")
	}

	if levelsMap["2"] {
		pnlMonitor2 := PNLMonitor{
			Logger:    myLogger,
			ApiKey:    apiKey,
			SecretKey: secretKey,
			WP:        wg,
			Cache:     cache,
			Filter:    FilterMap[FilterLevel2],
		}
		pnlMonitor2.Init()
		wg.Add(1)
		go pnlMonitor2.Run(shutdownCtx)
		myLogger.Debug("Enable Level2 report")
	}

	if levelsMap["3"] {
		pnlMonitor3 := PNLMonitor{
			Logger:    myLogger,
			ApiKey:    apiKey,
			SecretKey: secretKey,
			WP:        wg,
			Cache:     cache,
			Filter:    FilterMap[FilterLevel3],
		}
		pnlMonitor3.Init()
		wg.Add(1)
		go pnlMonitor3.Run(shutdownCtx)
		myLogger.Debug("Enable Level3 report")
	}

	if levelsMap["4"] {
		pnlMonitor4 := PNLMonitor{
			Logger:    myLogger,
			ApiKey:    apiKey,
			SecretKey: secretKey,
			WP:        wg,
			Cache:     cache,
			Filter:    FilterMap[FilterLevel4],
		}
		pnlMonitor4.Init()
		wg.Add(1)
		go pnlMonitor4.Run(shutdownCtx)
		myLogger.Debug("Enable Level4 report")
	}

	if levelsMap["5"] {
		pnlMonitor5 := PNLMonitor{
			Logger:    myLogger,
			ApiKey:    apiKey,
			SecretKey: secretKey,
			WP:        wg,
			Cache:     cache,
			Filter:    FilterMap[FilterLevel5],
		}
		pnlMonitor5.Init()
		wg.Add(1)
		go pnlMonitor5.Run(shutdownCtx)
		myLogger.Debug("Enable Level5 report")
	}

	if levelsMap["6"] {
		pnlMonitor6 := PNLMonitor{
			Logger:    myLogger,
			ApiKey:    apiKey,
			SecretKey: secretKey,
			WP:        wg,
			Cache:     cache,

			Filter: FilterMap[FilterLevel6],
		}
		pnlMonitor6.Init()
		wg.Add(1)
		go pnlMonitor6.Run(shutdownCtx)
		myLogger.Debug("Enable Level6 report")
	}

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

	wg.Wait()
}
