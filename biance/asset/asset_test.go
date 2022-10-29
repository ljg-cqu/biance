package asset

import (
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestGetUserAsset(t *testing.T) {
	client := &http.Client{}
	apiKey := ""
	secretKey := ""
	userAssets, err := GetAsset(client, biance.URLs[biance.URLUserAsset], "", apiKey, secretKey)
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
	userAssets, err := GetAsset(client, biance.URLs[biance.URLUserAsset], "BTC", apiKey, secretKey)
	require.Nil(t, err)
	fmt.Println(len(userAssets))
	for _, userAsset := range userAssets {
		fmt.Printf("%+v\n", userAsset)
	}
}
