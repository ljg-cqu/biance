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

func MappingBUSD(client biance.Client, priceUrl string, reverseFromBUSD bool, paris ...*Pair) ([]Mapped, error) {
	var editPair = paris
	var prices []price.Price
	for i, pair := range editPair {
		if !strings.Contains(string(pair.Symbol), "BUSD") && !strings.Contains(string(pair.Symbol), "USDT") {
			return nil, errors.New("only BUSD or USDT supported")
		}

		var pricesOne []price.Price
		var symbol = pair.Symbol
		var err error
		switch {
		case strings.HasSuffix(string(pair.Symbol), "BUSD"):
			pricesOne, err = getPriceFn(client, priceUrl, pair.Symbol)
			if err != nil {
				symbol = price.Symbol(strings.TrimSuffix(string(pair.Symbol), "BUSD") + "USDT")
				pricesOne, err = getPriceFn(client, priceUrl, symbol)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to get price for symbol %v", symbol)
				}
			}
		case strings.HasSuffix(string(pair.Symbol), "USDT"):
			pricesOne, err = getPriceFn(client, priceUrl, pair.Symbol)
			if err != nil {
				symbol = price.Symbol(strings.TrimSuffix(string(pair.Symbol), "USDT") + "BUSD")
				pricesOne, err = getPriceFn(client, priceUrl, symbol)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to get price for symbol %v", symbol)
				}
			}
		}

		editPair[i].Symbol = symbol
		prices = append(prices, pricesOne[0])
	}
	return mapping(prices, reverseFromBUSD, paris...), nil
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

func mapping(prices []price.Price, reverse bool, pairs ...*Pair) []Mapped {
	var pricesMap = make(map[price.Symbol]*big.Float)
	for _, price := range prices {
		pricesMap[price.Symbol] = price.Price
	}

	var mappeds []Mapped
	for _, pair := range pairs {
		price := pricesMap[pair.Symbol]

		var amtMapped *big.Float
		if reverse {
			amtMapped = new(big.Float).Quo(pair.BaseAmt, price)
		} else {
			amtMapped = new(big.Float).Mul(price, pair.BaseAmt)
		}

		mappeds = append(mappeds, Mapped{pair.Symbol, amtMapped, pair.BaseAmt, price})
	}
	return mappeds
}
