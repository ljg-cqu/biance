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

const (
	TrendUpStrengthen = iota
	TrendUpSteady
	TrendUpShake
	TrendUpWeaken

	TrendDownStrengthen
	TrendDownSteady
	TrendDownShake
	TrendDownWeaken

	TrendShake

	TrendZero
)

type Trend int

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

				s := t.Slices[token]
				s.AddElem(price_)
			}

			trends := t.trends(t.Slices)
			var trendTokensMap = make(map[Trend][]string)
			for token, trend := range trends {
				trendTokensMap[trend] = append(trendTokensMap[trend], string(token))
			}

			var trendTokensStrMap = make(map[Trend]string)
			for trend, tokens := range trendTokensMap {
				trendTokensStrMap[trend] = fmt.Sprintf("[%v] %v", len(tokens), strings.Join(tokens, ","))
			}

			var trendUpMakert, trendDownMarket, trendShakeMarket, trendZeroMarket string

			for trend, tokensStr := range trendTokensStrMap {
				switch trend {
				case TrendUpStrengthen:
					trendUpMakert = "++++++++++/" + "  " + tokensStr
				case TrendUpSteady:
					trendUpMakert = "++++++++++=" + "  " + tokensStr
				case TrendUpShake:
					trendUpMakert = "++++++++++~" + "  " + tokensStr
				case TrendUpWeaken:
					trendUpMakert = "++++++++++\\" + "  " + tokensStr
				case TrendDownStrengthen:
					trendDownMarket = "                ----------\\" + "  " + tokensStr
				case TrendDownSteady:
					trendDownMarket = "                ----------=" + "  " + tokensStr
				case TrendDownShake:
					trendDownMarket = "                ----------~" + "  " + tokensStr
				case TrendDownWeaken:
					trendDownMarket = "                ----------/" + "  " + tokensStr
				case TrendShake:
					trendShakeMarket = "                                +-+-+-+-+-" + "  " + tokensStr
				default:
					trendZeroMarket = "                                             0000000000" + "  " + tokensStr
				}
			}

			fmt.Printf("%v\n %v\n %v\n %v\n", trendUpMakert, trendDownMarket, trendShakeMarket, trendZeroMarket)
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

func (t *Trends) trends(slices map[price.Token]*slice.Slice) map[price.Token]Trend {
	if len(slices) == 0 {
		return nil
	}

	var trends_ = make(map[price.Token]Trend)

	for token, slice := range slices {
		if slice.Len() < t.PricesCountToMarkMicroTrend {
			continue
		}

		trend := t.trend(slice)
		trends_[token] = trend
	}

	return trends_
}

func (t *Trends) trend(s *slice.Slice) Trend {
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

	if positives == t.PricesCountToMarkMicroTrend-1 {
		var strenghthenNum, steadyNum, weakenNum int
		for i := 0; i < t.PricesCountToMarkMicroTrend-2; i++ {
			priceDifff := new(big.Float).Sub(priceDiffs[i+1], priceDiffs[i])
			switch priceDifff.Sign() {
			case -1:
				weakenNum++
			case 0:
				steadyNum++
			case 1:
				strenghthenNum++
			}
		}

		if strenghthenNum == t.PricesCountToMarkMicroTrend-2 {
			return TrendUpStrengthen
		} else if steadyNum == t.PricesCountToMarkMicroTrend-2 {
			return TrendUpSteady
		} else if weakenNum == t.PricesCountToMarkMicroTrend-2 {
			return TrendUpWeaken
		} else {
			return TrendUpShake
		}
	} else if negatives == t.PricesCountToMarkMicroTrend-1 {
		var strenghthenNum, steadyNum, weakenNum int
		for i := 0; i < t.PricesCountToMarkMicroTrend-2; i++ {
			priceDifff := new(big.Float).Sub(priceDiffs[i+1], priceDiffs[i])
			switch priceDifff.Sign() {
			case -1:
				strenghthenNum++
			case 0:
				steadyNum++
			case 1:
				weakenNum++
			}
		}

		if strenghthenNum == t.PricesCountToMarkMicroTrend-2 {
			return TrendDownStrengthen
		} else if steadyNum == t.PricesCountToMarkMicroTrend-2 {
			return TrendDownSteady
		} else if weakenNum == t.PricesCountToMarkMicroTrend-2 {
			return TrendDownWeaken
		} else {
			return TrendDownShake
		}
	} else if zeros == t.PricesCountToMarkMicroTrend-1 {
		return TrendZero
	} else {
		return TrendShake
	}
}
