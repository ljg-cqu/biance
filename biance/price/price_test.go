package price

import (
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestGetPricePairUSDTOverBUSD(t *testing.T) {
	client := &http.Client{}
	symbolPrices, err := GetPricePairUSDTOverBUSD(client, biance.URLs[biance.URLSymbolPrice])
	require.Nil(t, err)
	fmt.Println(len(symbolPrices))
	for _, symbolPrice := range symbolPrices {
		fmt.Printf("%+v\n", symbolPrice)
	}
}

func TestGetPricePairUSDTOverBUSDWithGivenTokens(t *testing.T) {
	client := &http.Client{}
	symbolPrices, err := GetPricePairUSDTOverBUSD(client, biance.URLs[biance.URLSymbolPrice], "SSV")
	require.Nil(t, err)
	fmt.Println(len(symbolPrices))
	for _, symbolPrice := range symbolPrices {
		fmt.Printf("%+v\n", symbolPrice)
	}
}

func TestGetPricePairBUSDOverUSDT(t *testing.T) {
	client := &http.Client{}
	symbolPrices, err := GetPricePairBUSDOverUSDT(client, biance.URLs[biance.URLSymbolPrice])
	require.Nil(t, err)
	fmt.Println(len(symbolPrices))
	for _, symbolPrice := range symbolPrices {
		fmt.Printf("%+v\n", symbolPrice)
	}
}

func TestGetPricePairBUSDOverUSDTWithGivenTokens(t *testing.T) {
	client := &http.Client{}
	symbolPrices, err := GetPricePairBUSDOverUSDT(client, biance.URLs[biance.URLSymbolPrice], "MFT")
	require.Nil(t, err)
	fmt.Println(len(symbolPrices))
	for _, symbolPrice := range symbolPrices {
		fmt.Printf("%+v\n", symbolPrice)
	}
}

func TestGetSymbolPricePairUSDT(t *testing.T) {
	client := &http.Client{}
	symbolPrices, err := GetPricePairUSDT(client, biance.URLs[biance.URLSymbolPrice])
	require.Nil(t, err)
	fmt.Println(len(symbolPrices))
	for _, symbolPrice := range symbolPrices {
		fmt.Printf("%+v\n", symbolPrice)
	}
}

func TestGetSymbolPricePairUSDTWithGivenTokens(t *testing.T) {
	client := &http.Client{}
	symbolPrices, err := GetPricePairUSDT(client, biance.URLs[biance.URLSymbolPrice], "BTC", "ETH")
	require.Nil(t, err)
	fmt.Println(len(symbolPrices))
	for _, symbolPrice := range symbolPrices {
		fmt.Printf("%+v\n", symbolPrice)
	}
}

func TestGetSymbolPricePairBUSD(t *testing.T) {
	client := &http.Client{}
	symbolPrices, err := GetPricePairBUSD(client, biance.URLs[biance.URLSymbolPrice])
	require.Nil(t, err)
	fmt.Println(len(symbolPrices))
	for _, symbolPrice := range symbolPrices {
		fmt.Printf("%+v\n", symbolPrice)
	}
}

func TestGetSymbolPricePairBUSDWithGivenTokens(t *testing.T) {
	client := &http.Client{}
	symbolPrices, err := GetPricePairBUSD(client, biance.URLs[biance.URLSymbolPrice], "BTC", "ETH")
	require.Nil(t, err)
	fmt.Println(len(symbolPrices))
	for _, symbolPrice := range symbolPrices {
		fmt.Printf("%+v\n", symbolPrice)
	}
}

func TestGetSymbolPrice(t *testing.T) {
	client := &http.Client{}
	symbolPrices, err := GetPrice(client, biance.URLs[biance.URLSymbolPrice])
	require.Nil(t, err)
	fmt.Println(len(symbolPrices))
	for _, symbolPrice := range symbolPrices {
		fmt.Printf("%+v\n", symbolPrice)
	}
}

func TestGetSymbolPriceWithGivenSymbols(t *testing.T) {
	client := &http.Client{}
	symbolPrices, err := GetPrice(client, biance.URLs[biance.URLSymbolPrice], "TRUBUSD", "TRUUSDT")
	require.Nil(t, err)
	fmt.Println(len(symbolPrices))
	for _, symbolPrice := range symbolPrices {
		fmt.Printf("%+v\n", symbolPrice)
	}
}
