package main

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"time"
)

const (
	TokenBUSD Token = "BUSD"
	TokenUSDT Token = "USDT"
)

type Token string

type Price struct {
	Symbol     string
	Price      string
	PriceFloat *big.Float
	When       time.Time `json:"-"`
}

type Prices []Price

func (p Prices) Len() int {
	return len(p)
}

func (p Prices) Less(i, j int) bool {
	if p[i].PriceFloat.Cmp(p[j].PriceFloat) == -1 {
		return true
	}
	return false
}

func (p Prices) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Prices) Sort() {
	sort.Sort(p)
}

func (p Prices) String() string {
	var str string
	for i, price := range p {
		str += fmt.Sprintf("%v %v:%v\n", i, price.Symbol, price.Price)
	}
	str += fmt.Sprintf("------------------\n")
	return str
}

// ---

type PricesOfOneSymbol []Price

func (p PricesOfOneSymbol) AveragePrice() *big.Float {
	if len(p) == 0 {
		return nil
	}
	var sum = new(big.Float)
	for _, price := range p {
		sum = sum.Add(sum, price.PriceFloat)
	}

	length, _ := new(big.Float).SetString(strconv.Itoa(len(p)))
	return new(big.Float).Quo(sum, length)
}

func (p PricesOfOneSymbol) LowPrice() *big.Float {
	if len(p) == 0 {
		return nil
	}
	var low = p[0].PriceFloat
	for _, price := range p {
		if r := price.PriceFloat.Cmp(low); r == -1 {
			low = price.PriceFloat
		}
	}
	return low
}

func (p PricesOfOneSymbol) HighPrice() *big.Float {
	if len(p) == 0 {
		return nil
	}
	var high = p[0].PriceFloat
	for _, price := range p {
		if r := price.PriceFloat.Cmp(high); r == 1 {
			high = price.PriceFloat
		}
	}
	return high
}
