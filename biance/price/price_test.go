package price

import (
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/stretchr/testify/require"
	"math/big"
	"net/http"
	"testing"
	"time"
)

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

func TestTrackPrice(t *testing.T) {
	var symbol = Symbol("ETHBUSD")
	client := &http.Client{}
	tk := time.NewTicker(time.Second)
	var nextIndex int
	var prices = [2]*Price{}
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			symbolPrices, err := GetPrice(client, biance.URLs[biance.URLSymbolPrice], symbol)
			require.Nil(t, err)
			price := symbolPrices[symbol]
			if nextIndex == 0 {
				prices[0] = &price
				nextIndex = 1
			} else {
				prices[1] = &price
				nextIndex = 0
			}
			if prices[0] != nil && prices[1] != nil {
				var newPrice, oldPrice *Price
				if nextIndex == 0 {
					newPrice = prices[1]
					oldPrice = prices[0]
				} else {
					newPrice = prices[0]
					oldPrice = prices[1]
				}
				one, _ := new(big.Float).SetString("1.0")
				quo := new(big.Float).Quo(newPrice.Price, oldPrice.Price)
				diff := new(big.Float).Sub(quo, one)
				fmt.Printf("%v, %v, %v, %v\n", symbol, oldPrice.Price, newPrice.Price, diff)
			}
		}
	}
}
