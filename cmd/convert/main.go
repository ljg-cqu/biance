package main

import (
	"flag"
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/asset"
	"github.com/ljg-cqu/biance/biance/pnl"
	"github.com/ljg-cqu/biance/logger"
	"math/big"
	"net/http"
	"strings"
)

func main() {
	var apiKey = flag.String("apiKey", "", "Binance API key")
	var secretKey = flag.String("secretKey", "", "Binance secret key")
	var levels = flag.String("reportLevels", "0 1 2 3 4 5 6", "report levels")
	var reportGain = flag.Bool("reportGain", true, "report gain")
	var reportLoss = flag.Bool("reportLoss", true, "report loss")
	flag.Parse()

	fmt.Printf("report levels: %v\n", *levels)
	fmt.Printf("report gain: %v\n", *reportGain)
	fmt.Printf("report loss: %v\n", *reportLoss)

	levelsSplit := strings.Fields(*levels)
	var levelsMap = make(map[string]bool)
	for _, level := range levelsSplit {
		levelsMap[level] = true
	}

	// Create logger
	// TODO: use simplified logger
	logger.DevMode = true
	logger.UseConsoleEncoder = true
	myLogger := logger.Default()

	client := &http.Client{}
	assetURL := biance.URLs[biance.URLUserAsset]
	priceURL := biance.URLs[biance.URLSymbolPrice]
	freePNLs, err := pnl.CheckFreePNLWithUSDTOrBUSD(client, assetURL, priceURL, "", *apiKey, *secretKey)
	if err != nil {
		myLogger.ErrorOnError(err, "failed to check pnl")
		return
	}

	var freePNLsFilter pnl.FreePNLs

	var tokenFilerMap = map[asset.Token]string{
		"NBS":  "",
		"USDT": "",
		"BUSD": "",
		"VIDT": "",
	}

	for _, freePNL := range freePNLs {
		_, ok := tokenFilerMap[freePNL.Token]
		if ok {
			continue
		}
		freePNLsFilter = append(freePNLsFilter, freePNL)
	}

	var gainPNLs []pnl.FreePNL
	var lossPNLs []pnl.FreePNL
	zeroGain, _ := new(big.Float).SetString("0")
	zeroLoss, _ := new(big.Float).SetString("0")

	gailThreshold, _ := new(big.Float).SetString("0.05")
	lossThreshold, _ := new(big.Float).SetString("0.05")
	var totalGain = zeroGain
	var totalLoss = zeroLoss

	oneHundred, _ := new(big.Float).SetString("100")
	zero, _ := new(big.Float).SetString("0")
	for _, freePNL := range freePNLsFilter {
		switch freePNL.PNLPercent.Cmp(zero) {
		case 1:
			if freePNL.PNLPercent.Cmp(gailThreshold) == 1 {
				gainPNLs = append(gainPNLs, freePNL)
				totalGain = new(big.Float).Add(totalGain, freePNL.PNLValue)
			}
		case 0:
		case -1:
			if new(big.Float).Abs(freePNL.PNLPercent).Cmp(lossThreshold) == 1 {
				lossPNLs = append(lossPNLs, freePNL)
				totalLoss = new(big.Float).Add(totalLoss, freePNL.PNLValue)
			}
		}
	}

	var gainInfoStr = fmt.Sprintf("\n++++++++++++++++++++tokens:%v profit:%v++++++++++++++++++++\n",
		len(gainPNLs), totalGain)
	for i, gainPNL := range gainPNLs {
		gainInfoStr += fmt.Sprintf("(%v) %v %v => %v @%v%%\n", i, gainPNL.Symbol, gainPNL.PNLAmountConvertable,
			gainPNL.PNLValue, new(big.Float).Mul(oneHundred, gainPNL.PNLPercent))
	}

	var lossInfoStr = fmt.Sprintf("--------------------tokens:%v loss:%v--------------------\n",
		len(lossPNLs), totalLoss)
	for i, lossPNL := range lossPNLs {
		lossInfoStr += fmt.Sprintf("(%v) %v %v @%v%%\n", i, lossPNL.Symbol,
			lossPNL.PNLValue, new(big.Float).Mul(oneHundred, lossPNL.PNLPercent))
	}

	fmt.Println(gainInfoStr)
	fmt.Println(lossInfoStr)
}
