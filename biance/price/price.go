package price

import (
	"encoding/json"
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type Symbol string
type Token string

type Price struct {
	Symbol Symbol
	Price  *big.Float
}

type price struct {
	Symbol Symbol `json:"symbol"`
	Price  string `json:"price"`
}

func GetPricePairUSDTOverBUSD(client biance.Client, url string, tokens ...Token) (map[Symbol]Price, error) {
	prices, err := GetPrice(client, url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var pricesPairBUSDUSDT = make(map[Symbol]Price)
	for symbol, price := range prices {
		if strings.HasSuffix(string(symbol), "USDT") {
			pricesPairBUSDUSDT[symbol] = price
		}
	}
	for symbol, price := range prices {
		if strings.HasSuffix(string(symbol), "BUSD") {
			_, ok := pricesPairBUSDUSDT[Symbol(strings.TrimSuffix(string(symbol), "BUSD")+"USDT")]
			if ok {
				continue
			}
			pricesPairBUSDUSDT[symbol] = price
		}
	}

	if tokens == nil {
		return pricesPairBUSDUSDT, nil
	}

	var pricesPairUSDTBUSDQuery = make(map[Symbol]Price)
	for _, token := range tokens {
		symbol := Symbol(token + "USDT")
		price, okBUSD := pricesPairBUSDUSDT[symbol]
		if okBUSD {
			pricesPairUSDTBUSDQuery[symbol] = price
			continue
		}
		symbol = Symbol(token + "BUSD")
		price, okUSDT := pricesPairBUSDUSDT[symbol]
		if okUSDT {
			pricesPairUSDTBUSDQuery[symbol] = price
			continue
		}
		return nil, errors.Errorf("no price found for token %v", token)
	}

	return pricesPairUSDTBUSDQuery, nil
}

func GetPricePairBUSDOverUSDT(client biance.Client, url string, tokens ...Token) (map[Symbol]Price, error) {
	prices, err := GetPrice(client, url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var pricesPairBUSDUSDT = make(map[Symbol]Price)
	for symbol, price := range prices {
		if strings.HasSuffix(string(symbol), "BUSD") {
			pricesPairBUSDUSDT[symbol] = price
		}
	}
	for symbol, price := range prices {
		if strings.HasSuffix(string(symbol), "USDT") {
			_, ok := pricesPairBUSDUSDT[Symbol(strings.TrimSuffix(string(symbol), "USDT")+"BUSD")]
			if ok {
				continue
			}
			pricesPairBUSDUSDT[symbol] = price
		}
	}

	if tokens == nil {
		return pricesPairBUSDUSDT, nil
	}

	var pricesPairBUSDUSDTQuery = make(map[Symbol]Price)
	for _, token := range tokens {
		symbol := Symbol(token + "BUSD")
		price, okBUSD := pricesPairBUSDUSDT[symbol]
		if okBUSD {
			pricesPairBUSDUSDTQuery[symbol] = price
			continue
		}
		symbol = Symbol(token + "USDT")
		price, okUSDT := pricesPairBUSDUSDT[symbol]
		if okUSDT {
			pricesPairBUSDUSDTQuery[symbol] = price
			continue
		}
		return nil, errors.Errorf("no price found for token %v", token)
	}

	return pricesPairBUSDUSDTQuery, nil
}

func GetPricePairUSDT(client biance.Client, url string, tokens ...Token) (map[Symbol]Price, error) {
	var symbols []Symbol
	for _, token := range tokens {
		symbols = append(symbols, Symbol(token+"USDT"))
	}
	prices, err := GetPrice(client, url, symbols...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var pricesPairBUSD = make(map[Symbol]Price)
	for symbol, price := range prices {
		if strings.HasSuffix(string(symbol), "USDT") {
			pricesPairBUSD[symbol] = price
		}
	}
	return pricesPairBUSD, nil
}

func GetPricePairBUSD(client biance.Client, url string, tokens ...Token) (map[Symbol]Price, error) {
	var symbols []Symbol
	for _, token := range tokens {
		symbols = append(symbols, Symbol(token+"BUSD"))
	}
	prices, err := GetPrice(client, url, symbols...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var pricesPairBUSD = make(map[Symbol]Price)
	for symbol, price := range prices {
		if strings.HasSuffix(string(symbol), "BUSD") {
			pricesPairBUSD[symbol] = price
		}
	}
	return pricesPairBUSD, nil
}

func GetPrice(client biance.Client, url string, symbols ...Symbol) (map[Symbol]Price, error) {
	var params string
	if symbols != nil {
		var symbolParam = "["
		for _, symbol := range symbols {
			symbolParam += fmt.Sprintf("\"%v\",", symbol)
		}
		symbolParam = strings.TrimSuffix(symbolParam, ",")
		symbolParam += "]"
		params = fmt.Sprintf("symbols=%v", symbolParam)
	}

	var payload = strings.NewReader("")
	req, err := http.NewRequest(http.MethodGet, url+"?"+params, payload)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create request")
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("request failed for %v", resp.Status)
	}

	var symbolPrices []price
	priceBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read response body")
	}

	err = json.Unmarshal(priceBytes, &symbolPrices)
	if err != nil {
		return nil, errors.Wrapf(err, "failed parse symbol price")
	}

	var SymbolPrices []Price
	for _, symbolPrice := range symbolPrices {
		priceFloat, _ := new(big.Float).SetString(symbolPrice.Price)
		SymbolPrices = append(SymbolPrices, Price{symbolPrice.Symbol, priceFloat})
	}

	var pricesMap = make(map[Symbol]Price)
	for _, price := range SymbolPrices {
		pricesMap[price.Symbol] = price
	}

	return pricesMap, nil
}
