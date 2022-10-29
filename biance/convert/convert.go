package convert

import (
	"encoding/json"
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/utils"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type ConvertResp struct {
	TranId uint64 `json:"tranId"`
	Status string `json:"status"`
}

type Params struct {
	ClientTranId string
	Asset        string
	Amount       *big.Float
	TargetAsset  string
}

func Convert(client biance.Client, url, apiKey, secretKey string, input *Params) (*ConvertResp, error) {
	var params string
	params = fmt.Sprintf("clientTranId=%v&asset=%v&amount=%v&targetAsset=%v",
		input.ClientTranId, input.Asset, input.Amount, input.TargetAsset)
	params = utils.CalculateAndAppendSignature(params, secretKey)

	var payload = strings.NewReader("")
	req, err := http.NewRequest(http.MethodPost, url+"?"+params, payload)
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

	var convertResp ConvertResp
	respBody, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBody, &convertResp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse response params")
	}

	return &convertResp, nil
}
