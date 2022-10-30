package mapping

import (
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/price"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

type Mapped struct {
	Symbol    price.Symbol
	TargetAmt *big.Float
	BaseAmt   *big.Float
	Price     *big.Float
}

type Pair struct {
	Symbol  price.Symbol
	BaseAmt *big.Float
}

var getPriceFn = price.GetPrice

func MappingUSDTOrBUSD(client biance.Client, priceUrl string, reverse bool, pairs ...*Pair) ([]Mapped, error) {
	var prices = make(map[price.Symbol]price.Price)
	pricesAllMap, err := getPriceFn(client, priceUrl)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get prices")
	}

	for i, pair := range pairs {
		if pair.Symbol == "" {
			return nil, errors.Errorf("Symbol shouldn't be empty")
		}
		if !strings.Contains(string(pair.Symbol), "BUSD") && !strings.Contains(string(pair.Symbol), "USDT") {
			return nil, errors.New("only BUSD or USDT supported")
		}

		price_, ok := pricesAllMap[pair.Symbol]
		var newSymbol = pair.Symbol
		if !ok {
			switch {
			case strings.HasSuffix(string(pair.Symbol), "BUSD"):
				newSymbol = price.Symbol(strings.TrimSuffix(string(pair.Symbol), "BUSD") + "USDT")
			case strings.HasSuffix(string(pair.Symbol), "USDT"):
				newSymbol = price.Symbol(strings.TrimSuffix(string(pair.Symbol), "USDT") + "BUSD")
			}
		}

		if !ok {
			price_, ok = pricesAllMap[price.Symbol(newSymbol)]
			if !ok {
				return nil, errors.Errorf("no price found for %v", pair.Symbol)
			}
		}

		pairs[i].Symbol = newSymbol
		prices[newSymbol] = price_
	}

	return mapping(prices, reverse, pairs...), nil
}

func Mapping(client biance.Client, priceUrl string, reverse bool, paris ...*Pair) ([]Mapped, error) {
	var symbols []price.Symbol
	for _, pair := range paris {
		symbols = append(symbols, pair.Symbol)
	}
	prices, err := getPriceFn(client, priceUrl, symbols...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get price")
	}

	return mapping(prices, reverse, paris...), nil
}

func mapping(pricesMap map[price.Symbol]price.Price, reverse bool, pairs ...*Pair) []Mapped {
	var mappeds []Mapped
	for _, pair := range pairs {
		price_ := pricesMap[pair.Symbol]

		var amtMapped *big.Float
		if reverse {
			amtMapped = new(big.Float).Quo(pair.BaseAmt, price_.Price)
		} else {
			amtMapped = new(big.Float).Mul(price_.Price, pair.BaseAmt)
		}

		mappeds = append(mappeds, Mapped{pair.Symbol, amtMapped, pair.BaseAmt, price_.Price})
	}
	return mappeds
}
