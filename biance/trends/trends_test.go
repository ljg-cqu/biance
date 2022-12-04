package trends

import (
	"context"
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/price"
	"github.com/ljg-cqu/biance/logger"
	"github.com/ljg-cqu/biance/utils/slice"
	"github.com/stretchr/testify/require"
	"math/big"
	"net/http"
	"testing"
	"time"
)

func TestTrends_TrackPairBUSDOrUSDT(t *testing.T) {
	logger.DevMode = true
	logger.UseConsoleEncoder = true
	myLogger := logger.Default()

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)

	trends := Trends{
		Logger:                      myLogger,
		ShutDownCtx:                 ctx,
		IntervalToQueryPrice:        1,
		PricesCountToMarkMicroTrend: 3,
		CheckPriceBUSDOverUSDT:      true,
		Client:                      &http.Client{},
	}

	trends.TrackPairBUSDOrUSDT()
}

func TestTrends_TrackPairBUSDOrUSDT_WithGivenToken(t *testing.T) {
	logger.DevMode = true
	logger.UseConsoleEncoder = true
	myLogger := logger.Default()

	ctx, _ := context.WithTimeout(context.Background(), time.Minute)

	trends := Trends{
		Logger:                      myLogger,
		ShutDownCtx:                 ctx,
		IntervalToQueryPrice:        1,
		PricesCountToMarkMicroTrend: 3,
		TokensToTrackPrice:          map[price.Token]bool{"BTC": true},
		CheckPriceBUSDOverUSDT:      true,
		Client:                      &http.Client{},
	}

	slices := make(map[price.Token]*slice.Slice)
	for token, _ := range trends.TokensToTrackPrice {
		slices[token] = slice.New(trends.PricesCountToMarkMicroTrend)
	}

	trends.Slices = slices

	trends.TrackPairBUSDOrUSDT()
}

func TestTrackPrice(t *testing.T) {
	var symbol = price.Symbol("BTCBUSD")
	client := &http.Client{}
	tk := time.NewTicker(time.Second * 5)
	var nextIndex int
	var prices = [2]*price.Price{}
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			symbolPrices, err := price.GetPrice(client, biance.URLs[biance.URLSymbolPrice], symbol)
			require.Nil(t, err)
			price_ := symbolPrices[symbol]
			if nextIndex == 0 {
				prices[0] = &price_
				nextIndex = 1
			} else {
				prices[1] = &price_
				nextIndex = 0
			}
			if prices[0] != nil && prices[1] != nil {
				var newPrice, oldPrice *price.Price
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
	var symbol = price.Symbol("BTCBUSD")
	client := &http.Client{}

	tk := time.NewTicker(time.Second * 1)
	intervalsToTrack := 3
	s := slice.New(intervalsToTrack)
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			symbolPrices, err := price.GetPrice(client, biance.URLs[biance.URLSymbolPrice], symbol)
			require.Nil(t, err)
			price_ := symbolPrices[symbol]
			s.AddElem(price_)
			if s.Len() < intervalsToTrack {
				continue
			}

			var priceI, priceJ *big.Float
			var prices = make([]*big.Float, intervalsToTrack)
			var priceDiffs = make([]*big.Float, intervalsToTrack-1)

			for i := 0; i < intervalsToTrack; i++ {
				priceI = s.Elem(i).(price.Price).Price
				prices[i] = priceI
				if i < intervalsToTrack-1 {
					priceJ = s.Elem(i + 1).(price.Price).Price
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

			if positives == intervalsToTrack-1 {
				market = "++++++++++"
			} else if negatives == intervalsToTrack-1 {
				market = "              ----------"
			} else if zeros == intervalsToTrack-1 {
				market = "                                         0000000000"
			} else {
				market = "                            +-+-+-+-+-"
			}

			switch priceDiffs[len(priceDiffs)-1].Sign() {
			case -1:
				suffix = " -"
			case 0:
				suffix = " 0"
			case 1:
				suffix = " +"
			}

			market += suffix
			fmt.Printf("%v: %v\n", symbol, market)
		}
	}
}
