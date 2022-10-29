package asset

import (
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestGetAssetWithDolVal(t *testing.T) {
	client := &http.Client{}
	apiKey := ""
	secretKey := ""
	assetURL := biance.URLs[biance.URLUserAsset]
	priceURL := biance.URLs[biance.URLSymbolPrice]
	userAssets, err := GetAssetWithDollar(client, assetURL, priceURL, "", apiKey, secretKey)
	require.Nil(t, err)
	fmt.Println(len(userAssets))
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
