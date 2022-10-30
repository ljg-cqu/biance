package pnl

import (
	"fmt"
	"github.com/ljg-cqu/biance/biance"
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

	var gainPNLs []FreePNL
	zero, _ := new(big.Float).SetString("0")
	var totalGain = zero
	for _, freePNL := range freePNLs {
		one, _ := new(big.Float).SetString("1")
		if freePNL.PNLValue.Cmp(one) == 1 {
			gainPNLs = append(gainPNLs, freePNL)
			totalGain = new(big.Float).Add(totalGain, freePNL.PNLValue)
		}
	}

	fmt.Printf("Total gain: %v tokens, about %v dollars\n\n", len(gainPNLs), totalGain)
	var gainToConvertStr string
	for _, gainPNL := range gainPNLs {
		oneHundred, _ := new(big.Float).SetString("100")
		gainToConvertStr += fmt.Sprintf("%v: %v => %v @%v%%\n", gainPNL.Symbol, gainPNL.PNLAmountConvertable,
			gainPNL.PNLValue, new(big.Float).Mul(oneHundred, gainPNL.PNLPercent))
	}
	fmt.Println(gainToConvertStr)

	fmt.Println("--------------Details------------------")
	fmt.Println(len(freePNLs))
	for _, freePNL := range freePNLs {
		fmt.Println(freePNL)
	}
}
