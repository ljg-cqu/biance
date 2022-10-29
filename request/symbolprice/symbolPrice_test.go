package symbolprice

import (
	"fmt"
	"github.com/ljg-cqu/biance/request"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestGetSymbolPrice(t *testing.T) {
	client := &http.Client{}
	symbolPrices, err := GetSymbolPrice(client, request.URLs[request.URLSymbolPrice])
	require.Nil(t, err)
	fmt.Println(len(symbolPrices))
	for _, symbolPrice := range symbolPrices {
		fmt.Printf("%+v\n", symbolPrice)
	}
}

func TestGetSymbolPriceWithGivenSymbols(t *testing.T) {
	client := &http.Client{}
	symbolPrices, err := GetSymbolPrice(client, request.URLs[request.URLSymbolPrice], "BTCUSDT", "BNBUSDT")
	require.Nil(t, err)
	fmt.Println(len(symbolPrices))
	for _, symbolPrice := range symbolPrices {
		fmt.Printf("%+v\n", symbolPrice)
	}
}
