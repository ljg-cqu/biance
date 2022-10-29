package biance

const (
	ApiKeyHeader = "X-MBX-APIKEY"
)

const (
	URLSymbolPrice URL = iota
	URLUserAsset
	URLConvertTransfer
)

const (
	endpoint = "https://api.binance.com"

	tickerPriceUrlPath = "/api/v3/ticker/price"
	userAssetUrlPath   = "/sapi/v3/asset/getUserAsset"
	convertTransfer    = "/sapi/v1/asset/convert-transfer"
)

type URL int

var URLs = map[URL]string{
	URLSymbolPrice:     endpoint + tickerPriceUrlPath,
	URLUserAsset:       endpoint + userAssetUrlPath,
	URLConvertTransfer: endpoint + convertTransfer,
}
