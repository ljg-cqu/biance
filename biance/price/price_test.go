package price

import (
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/utils/slice"
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
	var symbol = Symbol("BTCBUSD")
	client := &http.Client{}
	tk := time.NewTicker(time.Second * 5)
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

func TestTrackPrices(t *testing.T) {
	var symbol = Symbol("BTCBUSD")
	client := &http.Client{}

	tk := time.NewTicker(time.Second * 1)
	intervalsToTrack := 10
	s := slice.New(intervalsToTrack)
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			symbolPrices, err := GetPrice(client, biance.URLs[biance.URLSymbolPrice], symbol)
			require.Nil(t, err)
			price := symbolPrices[symbol]
			s.AddElem(price)
			if s.Len() < intervalsToTrack {
				continue
			}

			var priceI, priceJ *big.Float
			var prices = make([]*big.Float, intervalsToTrack)
			var priceDiffs = make([]*big.Float, intervalsToTrack-1)

			for i := 0; i < intervalsToTrack; i++ {
				priceI = s.Elem(i).(Price).Price
				prices[i] = priceI
				if i < intervalsToTrack-1 {
					priceJ = s.Elem(i + 1).(Price).Price
					priceDiff := new(big.Float).Sub(priceJ, priceI)
					priceDiffs[i] = priceDiff
				}
			}

			var negatives, zeros, positives int
			for _, priceDiff := range priceDiffs {
				switch priceDiff.Sign() {
				case -1:
					negatives++
				case 0:
					zeros++
				case 1:
					positives++
				}

			}

			var priceStr string
			for _, price := range prices {
				priceStr += fmt.Sprintf("%v,", price.Text('f', 10))
			}

			var priceDiffStr string
			for _, priceDiff := range priceDiffs {
				priceDiffStr += fmt.Sprintf("%v,", priceDiff.Text('f', 10))
			}

			//fmt.Printf("%v:\n %v\n %v\n", symbol, priceStr, priceDiffStr)

			var market string
			var suffix string

			if negatives == intervalsToTrack-1 {
				market = "----------"
			} else if zeros == intervalsToTrack-1 {
				market = "0000000000"
			} else if positives == intervalsToTrack-1 {
				market = "++++++++++"
			} else {
				market = "+-+-+-+-+-"
			}

			switch priceDiffs[len(priceDiffs)-1].Sign() {
			case -1:
				suffix = "     -"
			case 0:
				suffix = "     0"
			case 1:
				suffix = "     +"
			}

			market += suffix
			fmt.Sprintf("%v: %v\n", symbol, market)
		}
	}
}
