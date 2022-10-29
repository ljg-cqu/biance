package symbolprice

import (
	"encoding/json"
	"fmt"
	"github.com/ljg-cqu/biance/client"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type Symbol string

type SymbolPrice struct {
	Symbol Symbol
	Price  *big.Float
}

type symbolPrice struct {
	Symbol Symbol `json:"symbol"`
	Price  string `json:"price"`
}

func GetSymbolPrice(client client.Client, url string, symbols ...Symbol) ([]SymbolPrice, error) {
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

	var symbolPrices []symbolPrice
	priceBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read response body")
	}

	err = json.Unmarshal(priceBytes, &symbolPrices)
	if err != nil {
		return nil, errors.Wrapf(err, "failed parse symbol price")
	}

	var SymbolPrices []SymbolPrice
	for _, symbolPrice := range symbolPrices {
		priceFloat, _ := new(big.Float).SetString(symbolPrice.Price)
		SymbolPrices = append(SymbolPrices, SymbolPrice{symbolPrice.Symbol, priceFloat})
	}
	return SymbolPrices, nil
}
