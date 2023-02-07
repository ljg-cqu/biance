package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/asset"
	"github.com/ljg-cqu/biance/biance/pnl"
	"github.com/ljg-cqu/biance/logger"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

const (
	FileGainConvertFrom  = "gainConvertFrom.txt"
	FileGainConvertTo    = "gainConvertTo.txt"
	FileGainConvertValue = "gainConvertValue.txt"

	FileLossConvertFrom  = "lossConvertFrom.txt"
	FileLossConvertTo    = "lossConvertTo.txt"
	FileLossConvertValue = "lossConvertValue.txt"
)

func main() {
	var apiKey = flag.String("apiKey", "", "Binance API key")
	var secretKey = flag.String("secretKey", "", "Binance secret key")
	var levels = flag.String("reportLevels", "0 1 2 3 4 5 6", "report levels")
	var reportGain = flag.Bool("reportGain", true, "report gain")
	var reportLoss = flag.Bool("reportLoss", true, "report loss")
	var gainConvertT = flag.String("gainConvertT", "0.011", "gain convert threshold")
	var lossConvertT = flag.String("lossConvertT", "0.10", "loss convert threshold")
	var checkPNLInterval = flag.Int64("checkPNLInterval", 1000, "check PNL interval in milli-second")
	var excludeToken = flag.Bool("excludeToken", true, "exclude token to auto convert")
	var configPath = flag.String("configPath", "", "config path")
	var convertedPath = flag.String("convertedPath", "", "converted path")
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

	pnl.Principal200ConfigPath = fmt.Sprintf("%s/principal200.json", *configPath)

	interval := *checkPNLInterval
	ticker := time.NewTicker(time.Millisecond * time.Duration(interval))
	defer ticker.Stop()

	// Create fs watcher.
	watcher, err := fsnotify.NewWatcher()
	myLogger.FatalOnError(err, "failed to create fs watcher")
	defer watcher.Close()
	err = watcher.Add(*convertedPath)
	myLogger.FatalOnError(err, "failed to add path for fs watcher")
	if err != nil {
		log.Fatal(err)
	}

	shutdownCtx, shutdown := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-shutdownCtx.Done():
				return
			case event, ok := <-watcher.Events:
				myLogger.DebugOnTrue(!ok, "fs watcher event not ok")
				if event.Has(fsnotify.Write) {
					convertedToken, err := os.ReadFile(event.Name)
					if err != nil {
						myLogger.ErrorOnError(err, "failed to read file")
						continue
					}
					switch filepath.Base(event.Name) {
					case "gainConverted.txt":
						pendingToken, err := os.ReadFile(FileGainConvertFrom)
						if err != nil {
							myLogger.ErrorOnError(err, "failed to read file")
							continue
						}
						if string(convertedToken) == string(pendingToken) {
							writeFile(FileGainConvertFrom, "")
							writeFile(FileGainConvertTo, "")
							writeFile(FileGainConvertValue, "")
						}
					case "lossConverted.txt":
						pendingToken, err := os.ReadFile(FileLossConvertTo)
						if err != nil {
							myLogger.ErrorOnError(err, "failed to read file")
							continue
						}
						if string(convertedToken) == string(pendingToken) {
							writeFile(FileLossConvertTo, "")
							writeFile(FileLossConvertFrom, "")
							writeFile(FileLossConvertValue, "")
						}
					default:
						myLogger.Error("wrong converted path provided")
					}
				}
			case err, ok := <-watcher.Errors:
				myLogger.DebugOnTrue(!ok, "fs watcher event not ok")
				myLogger.ErrorOnError(err, "fs watcher got error")
			case <-ticker.C:
				freePNLs, err := pnl.CheckFreePNLWithUSDTOrBUSD(client, assetURL, priceURL, "", *apiKey, *secretKey)
				if err != nil {
					myLogger.ErrorOnError(err, "failed to check pnl")
					continue
				}

				var freePNLsFilter = freePNLs

				if *excludeToken {
					freePNLsFilter = nil
					// TODO: config infile and watch it for changes
					var tokenFilerMap = map[asset.Token]string{
						"NBS":  "",
						"USDT": "",
						"BUSD": "",
						"HIFI": "",
						"GFT":  "",
						//"VIDT": "",
						"AST": "",
					}

					for _, freePNL := range freePNLs {
						_, ok := tokenFilerMap[freePNL.Token]
						if ok {
							continue
						}
						freePNLsFilter = append(freePNLsFilter, freePNL)
					}
				}

				var gainPNLs []pnl.FreePNL
				var lossPNLs []pnl.FreePNL
				zeroGain, _ := new(big.Float).SetString("0")
				zeroLoss, _ := new(big.Float).SetString("0")

				// TODO: config it
				gailThreshold, _ := new(big.Float).SetString(*gainConvertT)
				lossThreshold, _ := new(big.Float).SetString(*lossConvertT)
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

				if *reportGain {
					err := writeFile("gain.txt", gainInfoStr)
					myLogger.ErrorOnError(err, "failed to write file")
					fmt.Println(gainInfoStr)

					if len(gainPNLs) > 0 {
						gainMax := gainPNLs[0]
						token := string(gainMax.Token)
						symbol := string(gainMax.Symbol)

						// TODO: deal with data consistency
						err := writeFile(FileGainConvertFrom, token)
						myLogger.ErrorOnError(err, "failed to write file")

						err = writeFile(FileGainConvertTo, strings.TrimPrefix(symbol, token))
						myLogger.ErrorOnError(err, "failed to write file")

						err = writeFile(FileGainConvertValue, gainMax.PNLAmountConvertable.Text('f', 10))
						myLogger.ErrorOnError(err, "failed to write file")
					} else {
						writeFile(FileGainConvertFrom, "")
						writeFile(FileGainConvertTo, "")
						writeFile(FileGainConvertValue, "")
					}
				}

				if *reportLoss {
					err := writeFile("loss.txt", lossInfoStr)
					myLogger.ErrorOnError(err, "failed to write file")
					fmt.Println(lossInfoStr)

					if len(lossPNLs) > 0 { // TODO: emit when have enough balance
						lossMax := lossPNLs[len(lossPNLs)-1]
						token := string(lossMax.Token)
						symbol := string(lossMax.Symbol)

						// TODO: deal with data consistency
						err := writeFile(FileLossConvertTo, token)
						myLogger.ErrorOnError(err, "failed to write file")

						err = writeFile(FileLossConvertFrom, strings.TrimPrefix(symbol, token))
						myLogger.ErrorOnError(err, "failed to write file")

						err = writeFile(FileLossConvertValue, strings.TrimPrefix(lossMax.PNLValue.Text('f', 10), "-"))
						myLogger.ErrorOnError(err, "failed to write file")
					} else {
						writeFile(FileLossConvertTo, "")
						writeFile(FileLossConvertFrom, "")
						writeFile(FileLossConvertValue, "")
					}
				}
			}
		}
	}()

	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		shutdown()
	}()
	<-shutdownCtx.Done()
}

func writeFile(path, content string) error {
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	_, err = f.WriteString(content)
	err = f.Close() // TODO: not write error?
	return err
}
