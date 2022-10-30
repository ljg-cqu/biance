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

	var considerableGianPNLs []FreePNL
	var considerableLossPNLs []FreePNL
	zeroGain, _ := new(big.Float).SetString("0")
	zeroLoss, _ := new(big.Float).SetString("0")

	one, _ := new(big.Float).SetString("1")
	considerableLossPercent, _ := new(big.Float).SetString("0.03")
	var totalGain = zeroGain
	var totalLoss = zeroLoss

	oneHundred, _ := new(big.Float).SetString("100")

	for _, freePNL := range freePNLsFilter {
		if freePNL.PNLValue.Cmp(one) == 1 {
			considerableGianPNLs = append(considerableGianPNLs, freePNL)
			totalGain = new(big.Float).Add(totalGain, freePNL.PNLValue)
			continue
		}
		if new(big.Float).Abs(freePNL.PNLPercent).Cmp(considerableLossPercent) == 1 {
			considerableLossPNLs = append(considerableLossPNLs, freePNL)
			totalLoss = new(big.Float).Add(totalLoss, freePNL.PNLValue)
		}
	}

	var gainInfoStr = fmt.Sprintf("\n++++++++++++++++++++tokens:%v profit:%v++++++++++++++++++++\n",
		len(considerableGianPNLs), totalGain)
	for _, gainPNL := range considerableGianPNLs {
		gainInfoStr += fmt.Sprintf("%v %v => %v @%v%%\n", gainPNL.Symbol, gainPNL.PNLAmountConvertable,
			gainPNL.PNLValue, new(big.Float).Mul(oneHundred, gainPNL.PNLPercent))
	}

	var lossInfoStr = fmt.Sprintf("--------------------tokens:%v loss:%v--------------------\n",
		len(considerableLossPNLs), totalLoss)
	for _, lossPNL := range considerableLossPNLs {
		lossInfoStr += fmt.Sprintf("%v %v @%v%%\n", lossPNL.Symbol,
			lossPNL.PNLValue, new(big.Float).Mul(oneHundred, lossPNL.PNLPercent))
	}

	fmt.Println(gainInfoStr)
	fmt.Println(lossInfoStr)
}
