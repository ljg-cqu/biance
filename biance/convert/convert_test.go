package convert

import (
	"fmt"
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/utils"
	"github.com/stretchr/testify/require"
	"math/big"
	"net/http"
	"testing"
)

func TestConvert(t *testing.T) {
	client := &http.Client{}
	apiKey := ""
	secretKey := ""
	amnt, _ := new(big.Float).SetString("0.24")
	var params = Params{
		ClientTranId: utils.TranID(),
		Asset:        "BUSD",
		Amount:       amnt,
		TargetAsset:  "USDC",
	}
	resp, err := Convert(client, biance.URLs[biance.URLConvertTransfer], apiKey, secretKey, &params)
	require.Nil(t, err)
	fmt.Println(resp)
}
