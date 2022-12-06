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
	MicroTrendUpStrengthen MicroTrend = iota
	MicroTrendUpSteady
	MircoTrendUpShake
	MicroTrendUpWeaken

	MicroTrendDownStrengthen
	MicroTrendDownSteady
	MicroTrendDownShake
	MicroTrendDownWeaken

	MicroTrendShakeUp
	MicroTrendShakeZero
	MicroTrendShakeDown

	MicroTrendZero
)

const (
	MacroTrendUp MacroTrend = iota
	MacroTrendDown
	MacroTrendShake
	MacroTrendZero
)

type MicroTrend int

type MacroTrend int

func isMicroTrendUp(t MicroTrend) bool {
	return t == MicroTrendUpStrengthen || t == MicroTrendUpSteady || t == MircoTrendUpShake || t == MicroTrendUpWeaken
}

func isMicroTrendDown(t MicroTrend) bool {
	return t == MicroTrendDownStrengthen || t == MicroTrendDownSteady || t == MicroTrendDownShake || t == MicroTrendDownWeaken
}

func isMicroTrendShake(t MicroTrend) bool {
	return t == MicroTrendShakeUp || t == MicroTrendShakeZero || t == MicroTrendShakeDown
}

type Trends struct {
	Logger                           logger.Logger
	ShutDownCtx                      context.Context
	IntervalToQueryPrice             time.Duration // unit: second
	MicroTrendsCountToMarkMacroTrend int
	PricesCountToMarkMicroTrend      int
	TokensToTrackPrice               map[price.Token]bool
	Slices                           map[price.Token]*slice.Slice
	CheckPriceBUSDOverUSDT           bool
	Client                           *http.Client
	microTrendsMap                   map[price.Token]*slice.Slice
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

			microTrends := t.microTrends(t.Slices)
			if t.microTrendsMap == nil {
				t.microTrendsMap = make(map[price.Token]*slice.Slice)
			}
			for token, trend := range microTrends {
				_, ok := t.microTrendsMap[token]
				if !ok {
					t.microTrendsMap[token] = slice.New(t.MicroTrendsCountToMarkMacroTrend)
				}

				t.microTrendsMap[token].AddElem(trend)
			}

			var microTrendsMapMap = make(map[price.Token]map[MicroTrend]int)
			for token, microTrends := range t.microTrendsMap {
				if microTrends.Len() < t.MicroTrendsCountToMarkMacroTrend {
					continue
				}
				var m = make(map[MicroTrend]int)
				for i := 0; i < microTrends.Len(); i++ {
					microTrend := microTrends.Elem(i).(MicroTrend)
					m[microTrend]++
				}
				microTrendsMapMap[token] = m
			}

			var MacroTrendsMap = make(map[price.Token]MacroTrend)

			for token, microTrendsMap := range microTrendsMapMap {
				var microTrendUpCount, microTrendDownCount, microTrendShakeCount, microTrendZeroCount, totalCount int
				for microTrend, count := range microTrendsMap {
					totalCount += count
					if isMicroTrendUp(microTrend) {
						microTrendUpCount++
					} else if isMicroTrendDown(microTrend) {
						microTrendDownCount++
					} else if isMicroTrendShake(microTrend) {
						microTrendShakeCount++
					} else {
						microTrendZeroCount++
					}
				}

				if microTrendUpCount > totalCount/2 {
					MacroTrendsMap[token] = MacroTrendUp
				} else if microTrendDownCount > totalCount/2 {
					MacroTrendsMap[token] = MacroTrendDown
				} else if microTrendZeroCount > totalCount/2 {
					MacroTrendsMap[token] = MacroTrendZero
				} else {
					MacroTrendsMap[token] = MacroTrendShake
				}
			}

			var MacroTrendTokensMap = make(map[MacroTrend][]string)
			for token, macroTrend := range MacroTrendsMap {
				MacroTrendTokensMap[macroTrend] = append(MacroTrendTokensMap[macroTrend], string(token))
			}

			var macroTrendTokensStrMap = make(map[MacroTrend]string)
			for trend, tokens := range MacroTrendTokensMap {
				macroTrendTokensStrMap[trend] = fmt.Sprintf("[%v] %v", len(tokens), strings.Join(tokens, ","))
			}

			var macroTrendUpMakert, macroTrendDownMarket, macroTrendShakeMarket, macroTrendZeroMarket string

			for trend, tokensStr := range macroTrendTokensStrMap {
				switch trend {
				case MacroTrendUp:
					macroTrendUpMakert = "++++++++++" + "  " + tokensStr
				case MacroTrendDown:
					macroTrendDownMarket = "                ----------" + "  " + tokensStr
				case MacroTrendShake:
					macroTrendShakeMarket = "                                +-+-+-+-+-" + "  " + tokensStr
				default:
					macroTrendZeroMarket = "                                             0000000000" + "  " + tokensStr
				}
			}

			fmt.Printf("%v\nmacro trends:\n%v\n %v\n %v\n %v\n", time.Now(), macroTrendUpMakert, macroTrendDownMarket, macroTrendShakeMarket, macroTrendZeroMarket)
			// ---

			var microTrendTokensMap = make(map[MicroTrend][]string)
			for token, microTrend := range microTrends {
				microTrendTokensMap[microTrend] = append(microTrendTokensMap[microTrend], string(token))
			}

			var microTrendTokensStrMap = make(map[MicroTrend]string)
			for trend, tokens := range microTrendTokensMap {
				microTrendTokensStrMap[trend] = fmt.Sprintf("[%v] %v", len(tokens), strings.Join(tokens, ","))
			}

			var microTrendUpMakert, microTrendDownMarket, microTrendShakeMarket, microTrendZeroMarket string

			for trend, tokensStr := range microTrendTokensStrMap {
				switch trend {
				case MicroTrendUpStrengthen:
					microTrendUpMakert = "++++++++++/" + "  " + tokensStr
				case MicroTrendUpSteady:
					microTrendUpMakert = "++++++++++=" + "  " + tokensStr
				case MircoTrendUpShake:
					microTrendUpMakert = "++++++++++~" + "  " + tokensStr
				case MicroTrendUpWeaken:
					microTrendUpMakert = "++++++++++\\" + "  " + tokensStr
				case MicroTrendDownStrengthen:
					microTrendDownMarket = "                ----------\\" + "  " + tokensStr
				case MicroTrendDownSteady:
					microTrendDownMarket = "                ----------=" + "  " + tokensStr
				case MicroTrendDownShake:
					microTrendDownMarket = "                ----------~" + "  " + tokensStr
				case MicroTrendDownWeaken:
					microTrendDownMarket = "                ----------/" + "  " + tokensStr
				case MicroTrendShakeUp:
					microTrendShakeMarket = "                                +-+-+-+-+-/" + "  " + tokensStr
				case MicroTrendShakeDown:
					microTrendShakeMarket = "                                +-+-+-+-+-\\" + "  " + tokensStr
				case MicroTrendShakeZero:
					microTrendShakeMarket = "                                +-+-+-+-+-0" + "  " + tokensStr
				default:
					microTrendZeroMarket = "                                             0000000000" + "  " + tokensStr
				}
			}

			fmt.Printf("micro trends:\n%v\n %v\n %v\n %v\n", microTrendUpMakert, microTrendDownMarket, microTrendShakeMarket, microTrendZeroMarket)
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

func (t *Trends) microTrends(slices map[price.Token]*slice.Slice) map[price.Token]MicroTrend {
	if len(slices) == 0 {
		return nil
	}

	var trends_ = make(map[price.Token]MicroTrend)

	for token, slice := range slices {
		if slice.Len() < t.PricesCountToMarkMicroTrend {
			continue
		}

		trend := t.microTrend(slice)
		trends_[token] = trend
	}

	return trends_
}

func (t *Trends) microTrend(s *slice.Slice) MicroTrend {
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
			return MicroTrendUpStrengthen
		} else if steadyNum == t.PricesCountToMarkMicroTrend-2 {
			return MicroTrendUpSteady
		} else if weakenNum == t.PricesCountToMarkMicroTrend-2 {
			return MicroTrendUpWeaken
		} else {
			return MircoTrendUpShake
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
			return MicroTrendDownStrengthen
		} else if steadyNum == t.PricesCountToMarkMicroTrend-2 {
			return MicroTrendDownSteady
		} else if weakenNum == t.PricesCountToMarkMicroTrend-2 {
			return MicroTrendDownWeaken
		} else {
			return MicroTrendDownShake
		}
	} else if zeros == t.PricesCountToMarkMicroTrend-1 {
		return MicroTrendZero
	} else {
		priceDiff := priceDiffs[len(priceDiffs)-1]
		switch priceDiff.Sign() {
		case -1:
			return MicroTrendShakeDown
		case 0:
			return MicroTrendShakeZero
		case 1:
			//return MicroTrendShakeUp
		}
		return MicroTrendShakeUp
	}
}
