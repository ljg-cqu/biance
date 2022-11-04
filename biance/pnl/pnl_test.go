package pnl

import (
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/asset"
	"github.com/stretchr/testify/require"
	"math/big"
	"net/http"
	"testing"
)

func TestCheckFreePNLWithUSDTOrBUSD(t *testing.T) {
	client := &http.Client{}
	apiKey := ""
	secretKey := ""
	assetURL := biance.URLs[biance.URLUserAsset]
	priceURL := biance.URLs[biance.URLSymbolPrice]
	freePNLs, err := CheckFreePNLWithUSDTOrBUSD(client, assetURL, priceURL, "", apiKey, secretKey)
	require.Nil(t, err)

	var freePNLsFilter FreePNLs

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

	var gainPNLs []FreePNL
	var lossPNLs []FreePNL
	zeroGain, _ := new(big.Float).SetString("0")
	zeroLoss, _ := new(big.Float).SetString("0")

	gailThreshold, _ := new(big.Float).SetString("0.05")
	lossThreshold, _ := new(big.Float).SetString("0.10")
	var totalGain = zeroGain
	var totalLoss = zeroLoss

	oneHundred, _ := new(big.Float).SetString("100")

	for _, freePNL := range freePNLsFilter {
		if freePNL.PNLPercent.Cmp(gailThreshold) == 1 {
			gainPNLs = append(gainPNLs, freePNL)
			totalGain = new(big.Float).Add(totalGain, freePNL.PNLValue)
			continue
		}
		if new(big.Float).Abs(freePNL.PNLPercent).Cmp(lossThreshold) == 1 {
			lossPNLs = append(lossPNLs, freePNL)
			totalLoss = new(big.Float).Add(totalLoss, freePNL.PNLValue)
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
