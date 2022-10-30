package asset

import (
	"encoding/json"
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/price"
	"github.com/ljg-cqu/biance/biance/utils"
	utilsTime "github.com/ljg-cqu/biance/utils/time"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"sort"
	"strings"
)

type Token string

type Asset struct {
	Token       Token
	Free        *big.Float
	Locked      *big.Float
	Freeze      *big.Float
	Withdrawing *big.Float

	price.Price
	FreeValue *big.Float
}

type asset_ struct {
	Token       Token  `json:"asset"`
	Free        string `json:"free"`
	Locked      string `json:"locked"`
	Freeze      string `json:"freeze"`
	Withdrawing string `json:"withdrawing"`
}

type Assets []Asset

func (a Assets) Len() int {
	return len(a)
}

func (a Assets) Less(i, j int) bool {
	if a[i].FreeValue.Cmp(a[j].FreeValue) == 1 {
		return true
	}
	return false
}

func (a Assets) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (p Assets) Sort() {
	sort.Sort(p)
}

var getPriceFn = price.GetPrice

func GetAssetWithUSDTOrBUSDFreeValue(client biance.Client, assetURL, priceURL, asset, apiKey, secretKey string) (Assets, error) {
	assets, err := GetAsset(client, assetURL, asset, apiKey, secretKey)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get asset")
	}
	pricesMap, err := getPriceFn(client, priceURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to et price")
	}

	for i, asset := range assets {
		symbol := price.Symbol(asset.Token + "USDT")
		price_, ok := pricesMap[symbol]
		if !ok {
			symbol = price.Symbol(asset.Token + "BUSD")
			price_, ok = pricesMap[symbol]
		}
		if !ok {
			continue
		}

		assets[i].Price = price_
		assets[i].FreeValue = new(big.Float).Mul(price_.Price, asset.Free)
	}
	assets.Sort()
	return assets, nil
}

func GetAsset(client biance.Client, assetURL, asset, apiKey, secretKey string) (Assets, error) {
	var params string
	if asset != "" {
		params = fmt.Sprintf("asset=%v&timestamp=%v", asset, utilsTime.Timestamp())
	} else {
		params = fmt.Sprintf("timestamp=%v", utilsTime.Timestamp())
	}
	params = utils.CalculateAndAppendSignature(params, secretKey)

	var payload = strings.NewReader("")
	req, err := http.NewRequest(http.MethodPost, assetURL+"?"+params, payload)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create request")
	}
	req.Header.Set(biance.ApiKeyHeader, apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to post request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("request failed for %v", resp.Status)
	}

	var userAssets []asset_
	respBody, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBody, &userAssets)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse response params")
	}

	var UserAssets Assets

	for _, userAsset := range userAssets {
		free, _ := new(big.Float).SetString(userAsset.Free)
		locked, _ := new(big.Float).SetString(userAsset.Locked)
		freeze, _ := new(big.Float).SetString(userAsset.Freeze)
		withdrawing, _ := new(big.Float).SetString(userAsset.Withdrawing)

		var asset Asset
		asset.Token = userAsset.Token
		asset.Free = free
		asset.Locked = locked
		asset.Freeze = freeze
		asset.Withdrawing = withdrawing
		UserAssets = append(UserAssets, asset)
	}
	return UserAssets, nil
}
