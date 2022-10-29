package userdata

import (
	"encoding/json"
	"fmt"
	"github.com/ljg-cqu/biance/client"
	"github.com/ljg-cqu/biance/request"
	"github.com/ljg-cqu/biance/request/utils"
	utilsTime "github.com/ljg-cqu/biance/utils/time"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type UserAsset struct {
	Asset       string
	Free        *big.Float
	Locked      *big.Float
	Freeze      *big.Float
	Withdrawing *big.Float
}

type userAsset struct {
	Asset       string `json:"asset"`
	Free        string `json:"free"`
	Locked      string `json:"locked"`
	Freeze      string `json:"freeze"`
	Withdrawing string `json:"withdrawing"`
}

func getUserAsset(client client.Client, url, asset, apiKey, secretKey string) ([]UserAsset, error) {
	var body string
	if asset != "" {
		body = fmt.Sprintf("asset=%v&timestamp=%v", asset, utilsTime.Timestamp())
		body = utils.CalculateAndAppendSignature(body, secretKey)
	} else {
		body = fmt.Sprintf("timestamp=%v", utilsTime.Timestamp())
		body = utils.CalculateAndAppendSignature(body, secretKey)
	}

	var payload = strings.NewReader("")
	req, err := http.NewRequest(http.MethodPost, url+"?"+body, payload)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create request")
	}
	req.Header.Set(request.ApiKeyHeader, apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to post request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("request failed for %v", resp.Status)
	}

	var userAssets []userAsset
	respBody, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBody, &userAssets)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse response body")
	}

	var UserAssets []UserAsset

	for _, userAsset := range userAssets {
		free, _ := new(big.Float).SetString(userAsset.Free)
		locked, _ := new(big.Float).SetString(userAsset.Locked)
		freeze, _ := new(big.Float).SetString(userAsset.Freeze)
		withdrawing, _ := new(big.Float).SetString(userAsset.Withdrawing)
		UserAssets = append(UserAssets, UserAsset{userAsset.Asset, free, locked, freeze, withdrawing})
	}
	return UserAssets, nil
}
