package trends

import (
	"context"
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/price"
	"github.com/ljg-cqu/biance/logger"
	"github.com/ljg-cqu/biance/utils/slice"
	"github.com/pkg/errors"
	"math/big"
	"net/http"
	"strings"
	"time"
)

type Trends struct {
	Logger                      logger.Logger
	ShutDownCtx                 context.Context
	IntervalToQueryPrice        time.Duration // unit: second
	PricesCountToMarkMicroTrend int
	TokensToTrackPrice          map[price.Token]bool
	Slices                      map[price.Token]*slice.Slice
	CheckPriceBUSDOverUSDT      bool
	Client                      *http.Client
}

func (t *Trends) TrackPairBUSDOrUSDT() {
	tk := time.NewTicker(t.IntervalToQueryPrice * time.Second)
	defer tk.Stop()
	for {
		select {
		case <-t.ShutDownCtx.Done():
			return
		case <-tk.C:
			tokenPriceFoundMap, err := t.getAllPricesPairWithBUSDorUSDT()
			if err != nil {
				t.Logger.ErrorOnError(err, "failed to get all prices")
				continue
			}

			// todo: do once
			if t.TokensToTrackPrice == nil {
				t.TokensToTrackPrice = make(map[price.Token]bool)

				if t.Slices == nil {
					t.Slices = make(map[price.Token]*slice.Slice)
				}

				for token, _ := range tokenPriceFoundMap {
					t.TokensToTrackPrice[token] = true

					_, ok := t.Slices[token]
					if !ok {
						t.Slices[token] = slice.New(t.PricesCountToMarkMicroTrend)
					}
				}
			}

			for token, _ := range t.TokensToTrackPrice {
				price_, ok := tokenPriceFoundMap[token]
				if !ok {
					t.Logger.Error("not found price for token", []logger.Field{{"token", token}}...)
					continue
				}

				symbol := price_.Symbol
				s := t.Slices[token]
				s.AddElem(price_)

				if s.Len() < t.PricesCountToMarkMicroTrend {
					continue
				}

				var priceI, priceJ *big.Float
				var prices = make([]*big.Float, t.PricesCountToMarkMicroTrend)
				var priceDiffs = make([]*big.Float, t.PricesCountToMarkMicroTrend-1)

				for i := 0; i < t.PricesCountToMarkMicroTrend; i++ {
					priceI = s.Elem(i).(price.Price).Price
					prices[i] = priceI
					if i < t.PricesCountToMarkMicroTrend-1 {
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

				if positives == t.PricesCountToMarkMicroTrend-1 {
					market = "++++++++++"
				} else if negatives == t.PricesCountToMarkMicroTrend-1 {
					market = "              ----------"
				} else if zeros == t.PricesCountToMarkMicroTrend-1 {
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
}

func (t *Trends) getAllPricesPairWithBUSDorUSDT() (map[price.Token]price.Price, error) {
	var symbolPricesBUSDUSDT map[price.Symbol]price.Price
	var err error
	if t.CheckPriceBUSDOverUSDT {
		symbolPricesBUSDUSDT, err = price.GetPricePairBUSDOverUSDT(t.Client, biance.URLs[biance.URLSymbolPrice])
		if err != nil {
			return nil, errors.WithStack(err)
		}
	} else {
		symbolPricesBUSDUSDT, err = price.GetPricePairUSDTOverBUSD(t.Client, biance.URLs[biance.URLSymbolPrice])
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	var tokenPriceFoundMap = make(map[price.Token]price.Price)
	for symbol, price_ := range symbolPricesBUSDUSDT {
		symbolTim := strings.TrimSuffix(string(symbol), "BUSD")
		symbolTim = strings.TrimSuffix(symbolTim, "USDT")
		tokenPriceFoundMap[price.Token(symbolTim)] = price_
	}
	return tokenPriceFoundMap, nil
}
