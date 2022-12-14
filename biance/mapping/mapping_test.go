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

func TestMappingBUSD(t *testing.T) { // TODO: improve, reduce many request
	TokenValueRaw := "MDX:43 AKRO:9.111 SLP:8.95 HNT:13.47 JASMY:8.81 GALA:5.17"
	tokenValuePairs := strings.Fields(TokenValueRaw)

	var pairs []*Pair

	for _, tokenValuePair := range tokenValuePairs {
		tokenValue := strings.Split(tokenValuePair, ":")
		amtToMap, _ := new(big.Float).SetString(tokenValue[1])
		token := tokenValue[0]
		symbol := price.Symbol(token + "USDT")
		pair := Pair{
			Symbol:  symbol,
			BaseAmt: amtToMap,
		}
		pairs = append(pairs, &pair)
	}

	client := &http.Client{}
	mappeds, err := MappingUSDTOrBUSD(client, biance.URLs[biance.URLSymbolPrice], true, pairs...)
	require.Nil(t, err)
	for _, mapped := range mappeds {
		fmt.Printf("%+v\n", mapped)
	}
}

func TestMapping(t *testing.T) {
	TokenValueRaw := "EGLD:2.29832413"
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
	mappeds, err := Mapping(client, biance.URLs[biance.URLSymbolPrice], false, pairs...)
	require.Nil(t, err)
	for _, mapped := range mappeds {
		fmt.Printf("%+v\n", mapped)
	}
}
