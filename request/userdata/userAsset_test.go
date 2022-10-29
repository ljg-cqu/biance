package userdata

import (
	"fmt"
	"github.com/ljg-cqu/biance/request"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestGetUserAsset(t *testing.T) {
	client := &http.Client{}
	apiKey := "XesIIB7fra4VeCZQNovtnDgQI5QFt8073p3vWZBMIfbm3OLe0zl4A5YDs37961H0"
	secretKey := "GDpqvNk4spXqKfEQom9nr852NDmmaCxKWpDTdtoGZxUW2uGsl37rX5iMu9JBhBam"
	userAssets, err := getUserAsset(client, request.URLs[request.URLUserAsset], "", apiKey, secretKey)
	require.Nil(t, err)
	fmt.Println(len(userAssets))
	for _, userAsset := range userAssets {
		fmt.Printf("%+v\n", userAsset)
	}
}

func TestGetUserAssetWithGivenAsset(t *testing.T) {
	client := &http.Client{}
	apiKey := "XesIIB7fra4VeCZQNovtnDgQI5QFt8073p3vWZBMIfbm3OLe0zl4A5YDs37961H0"
	secretKey := "GDpqvNk4spXqKfEQom9nr852NDmmaCxKWpDTdtoGZxUW2uGsl37rX5iMu9JBhBam"
	userAssets, err := getUserAsset(client, request.URLs[request.URLUserAsset], "BTC", apiKey, secretKey)
	require.Nil(t, err)
	fmt.Println(len(userAssets))
	for _, userAsset := range userAssets {
		fmt.Printf("%+v\n", userAsset)
	}
}
