package asset

import (
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/stretchr/testify/require"
	"math/big"
	"net/http"
	"strings"
	"testing"
)

func TestGetAssetWithUSDTOrBUSDFreeValue(t *testing.T) {
	client := &http.Client{}
	apiKey := "f4e3ss8FqJdTgmJGQ5ZPfZa39x37eBi9MCrs8tfXtqwGY0CeBc1HA1YrZq6Au2fB"
	secretKey := "EE2WXi7v0cuOTsREu3CmYtLpJclLtGPVgKKGJLMMwDJ7xq97nvXZejjgoouMQzcE"
	assetURL := biance.URLs[biance.URLUserAsset]
	priceURL := biance.URLs[biance.URLSymbolPrice]
	userAssets, err := GetAssetWithUSDTOrBUSDFreeValue(client, assetURL, priceURL, "", apiKey, secretKey)
	require.Nil(t, err)
	fmt.Println(len(userAssets))
	principal, _ := new(big.Float).SetString("150")
	var richTokens string
	for _, userAsset := range userAssets {
		if userAsset.FreeValue.Cmp(principal) == 1 {
			richTokens += fmt.Sprintf("\"%v\",", userAsset.Token)
		}
	}
	fmt.Println(richTokens)
	for _, userAsset := range userAssets {
		token := string(userAsset.Symbol)
		token = strings.TrimSuffix(token, "USDT")
		token = strings.TrimSuffix(token, "BUSD")
		fmt.Printf("%v:%+v (%v)\n", token, userAsset.FreeValue, userAsset.Symbol)
	}
	fmt.Println("-----------------")
	for _, userAsset := range userAssets {
		fmt.Printf("%+v\n", userAsset)
	}
}

func TestGetUserAsset(t *testing.T) {
	client := &http.Client{}
	apiKey := ""
	secretKey := ""
	assetURL := biance.URLs[biance.URLUserAsset]
	//priceURL := biance.URLs[biance.URLSymbolPrice]
	userAssets, err := GetAsset(client, assetURL, "", apiKey, secretKey)
	require.Nil(t, err)
	fmt.Println(len(userAssets))
	for _, userAsset := range userAssets {
		fmt.Printf("%+v\n", userAsset)
	}
}

func TestGetUserAssetWithGivenAsset(t *testing.T) {
	client := &http.Client{}
	apiKey := ""
	secretKey := ""
	assetURL := biance.URLs[biance.URLUserAsset]
	//priceURL := biance.URLs[biance.URLSymbolPrice]
	userAssets, err := GetAsset(client, assetURL, "ETH", apiKey, secretKey)
	require.Nil(t, err)
	fmt.Println(len(userAssets))
	for _, userAsset := range userAssets {
		fmt.Printf("%+v\n", userAsset)
	}
}
