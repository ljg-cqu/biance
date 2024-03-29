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
	if len(assets) == 0 {
		return nil, errors.Errorf("found no asset")
	}

	pricesMap, err := getPriceFn(client, priceURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to et price")
	}

	if len(pricesMap) == 0 {
		return nil, errors.New("no price found")
	}

	var assetsToUSDTorBUSD Assets
	for i, asset := range assets {
		symbol := price.Symbol(asset.Token + "USDT")
		price_, ok := pricesMap[symbol]
		if !ok {
			symbol = price.Symbol(asset.Token + "BUSD")
			price_, ok = pricesMap[symbol]
		}
		if !ok && asset.Token != "USDT" && asset.Token != "BUSD" {
			fmt.Printf("Token %v has no paired price to USDT or BUSD\n", asset.Token)
			continue
		}

		if asset.Token == "USDT" || asset.Token == "BUSD" {
			assets[i].FreeValue = asset.Free
			one, _ := new(big.Float).SetString("1")
			var symbol price.Symbol
			if asset.Token == "USDT" {
				symbol = "USDTUSDT"
			} else {
				symbol = "BUSDBUSD"
			}
			assets[i].Price = price.Price{Symbol: symbol, Price: one}
			continue
		}

		assets[i].Price = price_
		assetsToUSDTorBUSD = append(assetsToUSDTorBUSD, assets[i])
	}

	for i, asset := range assetsToUSDTorBUSD {
		freeValue := new(big.Float).Mul(asset.Price.Price, asset.Free)
		//zero, _ := new(big.Float).SetString("0")
		//cmp := freeValue.Cmp(zero)
		//if cmp == 0 || cmp == -1 {
		//	fmt.Printf("Token %v has no free value to USDT or BUSD\n", asset.Token)
		//	continue
		//}
		assetsToUSDTorBUSD[i].FreeValue = freeValue
	}
	assetsToUSDTorBUSD.Sort()
	return assetsToUSDTorBUSD, nil
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
