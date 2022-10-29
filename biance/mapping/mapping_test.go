package mapping

import (
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/price"
	"github.com/stretchr/testify/require"
	"math/big"
	"net/http"
	"strings"
	"testing"
)

func TestMappingWithBUSDOrUSDT(t *testing.T) {
	TokenValueRaw := "KLAY:8.72 NEO:2.76 XEM:2.64 BTTC:2.48 POLS:2.21 BAT:2 GALA:1.97 " +
		"STORJ:1.95 LRC:1.92 BAL:1.9 API3:1.89 SKL:1.84"
	tokenValuePairs := strings.Fields(TokenValueRaw)

	var pairs []*Pair

	for _, tokenValuePair := range tokenValuePairs {
		tokenValue := strings.Split(tokenValuePair, ":")
		amtToMap, _ := new(big.Float).SetString(tokenValue[1])
		token := tokenValue[0]
		symbol := price.Symbol(token + "BUSD")
		pair := Pair{
			Symbol:  symbol,
			BaseAmt: amtToMap,
		}
		pairs = append(pairs, &pair)
	}

	client := &http.Client{}
	mappeds, err := MappingWithBUSDOrUSDT(client, biance.URLs[biance.URLSymbolPrice], true, pairs...)
	require.Nil(t, err)
	for _, mapped := range mappeds {
		fmt.Printf("%+v\n", mapped)
	}
}

func TestMapping(t *testing.T) {
	TokenValueRaw := "MDX:11.18 AERGO:4.02 OCEAN:3.38 PYR:3.21 BCH:2.49 COTI:2.47 WAVES:2.43 ALPINE:2.42 WTC:2.38 " +
		"FLM:2.36 NKN:2.21 AMB:2.21 AXS:2.20 REEF:2.09 BTTC:2.08 ASR:2.08 GAL:2.06 CREAM:2.06 ETH:2.02"
	tokenValuePairs := strings.Fields(TokenValueRaw)

	var pairs []*Pair

	for _, tokenValuePair := range tokenValuePairs {
		tokenValue := strings.Split(tokenValuePair, ":")
		amtToMap, _ := new(big.Float).SetString(tokenValue[1])
		token := tokenValue[0]
		symbol := price.Symbol(token + "BUSD")
		if token == "WTC" {
			symbol = price.Symbol(token + "USDT")
		}
		pair := Pair{
			Symbol:  symbol,
			BaseAmt: amtToMap,
		}
		pairs = append(pairs, &pair)
	}

	client := &http.Client{}
	mappeds, err := Mapping(client, biance.URLs[biance.URLSymbolPrice], true, pairs...)
	require.Nil(t, err)
	for _, mapped := range mappeds {
		fmt.Printf("%+v\n", mapped)
	}
}
