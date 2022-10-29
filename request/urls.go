package request

const (
	URLTickerPrice URL = iota
	URLUserAsset
)

const (
	endpoint = "https://api.binance.com"

	tickerPriceUrlPath = "/api/v3/ticker/price"
	userAssetUrlPath   = "/sapi/v3/asset/getUserAsset"
)

type URL int

var URLs = map[URL]string{
	URLTickerPrice: endpoint + tickerPriceUrlPath,
	URLUserAsset:   endpoint + userAssetUrlPath,
}