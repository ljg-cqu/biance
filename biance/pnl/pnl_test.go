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

func TestExtractToken(t *testing.T) {
	str := `(0) WNXMUSDT 4.783728306400000005 @4.783728306400000005%
(1) POWRUSDT 4.748201440064000002 @4.748201440064000002%
(2) SUPERUSDT 4.6965317931890000064 @4.6965317931890000064%
(3) KLAYUSDT 4.6884893986160000035 @4.6884893986160000035%
(4) AMPUSDT 4.686285384251600003 @4.686285384251600003%
(6) STMXUSDT 4.641861947268349997 @4.641861947268349997%
(7) KAVAUSDT 4.528472491039999999 @4.528472491039999999%
(8) ALICEUSDT 4.5256270548499999973 @4.5256270548499999973%
(9) BNTUSDT 4.489795921919999995 @4.489795921919999995%
(10) EGLDUSDT 4.4752331542000000006 @4.4752331542000000006%
(15) SRMUSDT 4.3594399474999999955 @4.3594399474999999955%
(19) MIRUSDT 4.146875167123799999 @4.146875167123799999%
(21) MITHUSDT 4.135251190999999993 @4.135251190999999993%
(22) BEAMUSDT 4.1200706314950000034 @4.1200706314950000034%
(23) SCRTUSDT 4.0946890513399999967 @4.0946890513399999967%
(24) LRCUSDT 4.030910477403 @4.030910477403%
(26) NEOUSDT 3.9954338432000000014 @3.9954338432000000014%
(27) ONGUSDT 3.928136419894999995 @3.9281364198949999947%
(28) NMRUSDT 3.9241470088999999971 @3.9241470088999999971%
(29) BSWUSDT 3.9009670167049999984 @3.9009670167049999986%`

	strLines := strings.Split(str, "\n")
	var tokens string
	for _, strLine := range strLines {
		elems := strings.Fields(strLine)
		symbol := elems[1]
		suffixs := []string{"BUSD", "USDT"}
		for _, suffix := range suffixs {
			symbol = strings.TrimSuffix(symbol, suffix)
		}
		tokens += fmt.Sprintf("%q,", symbol)
	}
	fmt.Println(tokens)
}
