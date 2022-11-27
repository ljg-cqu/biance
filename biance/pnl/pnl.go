package pnl

import (
	"github.com/ljg-cqu/biance/biance"
	"github.com/ljg-cqu/biance/biance/asset"
	"github.com/ljg-cqu/biance/biance/price"
	"github.com/pkg/errors"
	"math/big"
	"sort"
)

var tokenWithDollarPrincipal = map[asset.Token]string{
	//"MDX":   "1200",
	//"APT":   "1000",
	//"POLYX": "1000",
}

var tokenWith200DollarPrincipal = []asset.Token{
	"NULS", "ASR", "MOB", "DF", "BOND", "VTHO", "ANKR", "MBOX", "ALPINE", "HIGH", "TKO",
	"ZEC", "PROS", "OMG", "UTK", "PEOPLE", "MINA", "SUSHI", "DYDX", "YGG", "YFII", "RAY",
	"APT", "POLYX", "REN", "VOXEL", "TVK", "SC", "CHESS", "REI", "AVAX", "SANTOS",
	"MATIC", "ATOM", "BETH", "CTSI", "HBAR", "AION", "GLMR", "JOE", "UFT", "PORTO", "SSV",
	"KDA", "CFX", "XEC", "IOTX", "SLP", "AKRO", "DATA", "SUN", "ICX", "DOT", "MBL", "MFT",
	"KEY", "WING", "IDEX", "USTC", "PERP", "RIF", "MC", "PLA", "OSMO", "OG", "RAD", "ZRX",
	"LOOM", "OXT", "CVP", "POND", "TRU", "CLV", "QI", "PERL", "MULTI", "ERN", "FIO",
	"REQ", "PNT", "HIVE", "FIL", "COCOS", "LTO", "DREP", "LDO", "PROM", "RLC", "AUCTION",
	"WNXM", "POWR", "SUPER", "KLAY", "AMP", "STMX", "KAVA", "BNT", "EGLD", "SRM", "BURGER",
	"MIR", "MITH", "BEAM", "SCRT", "LRC", "NEO", "ONG", "NMR", "BSW"}

type FreePNL struct {
	Token asset.Token
	Free  *big.Float
	price.Price
	PrincipalValueInDollar *big.Float
	FreeValue              *big.Float
	PNLPercent             *big.Float
	PNLValue               *big.Float
	PNLAmountConvertable   *big.Float
}

type FreePNLs []FreePNL

func (p FreePNLs) Len() int {
	return len(p)
}

func (p FreePNLs) Less(i, j int) bool {
	if p[i].PNLPercent.Cmp(p[j].PNLPercent) == 1 {
		return true
	}
	return false
}

func (p FreePNLs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p FreePNLs) Sort() {
	sort.Sort(p)
}

var assetFn = asset.GetAssetWithUSDTOrBUSDFreeValue

func CheckFreePNLWithUSDTOrBUSD(client biance.Client, assetURL, priceURL, asset, apiKey, secretKey string) (FreePNLs, error) {
	assets, err := assetFn(client, assetURL, priceURL, asset, apiKey, secretKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get asset")
	}

	if len(assets) == 0 {
		return nil, nil
	}

	var freePNLs FreePNLs
	for _, asset := range assets {
		freePNLs = append(freePNLs, FreePNL{
			Token:     asset.Token,
			Free:      asset.Free,
			Price:     asset.Price,
			FreeValue: asset.FreeValue,
		})
	}

	configPrincipalValue(freePNLs)

	for i, freePNL := range freePNLs {
		pnlVal := new(big.Float).Sub(freePNL.FreeValue, freePNL.PrincipalValueInDollar)
		freePNLs[i].PNLValue = pnlVal

		pnlPercent := new(big.Float).Quo(pnlVal, freePNL.PrincipalValueInDollar)
		freePNLs[i].PNLPercent = pnlPercent

		//if freePNL.FreeValue.Cmp(freePNL.PrincipalValueInDollar) == 1 {
		pnlAmtConvertable := new(big.Float).Quo(pnlVal, freePNL.Price.Price)
		freePNLs[i].PNLAmountConvertable = pnlAmtConvertable
		//} else {
		//	zero, _ := new(big.Float).SetString("0")
		//	freePNLs[i].PNLAmountConvertable = zero
		//}
	}

	freePNLs.Sort()
	return freePNLs, nil
}

// TODO:

func configPrincipalValue(pnls FreePNLs) {
	var token200Map = make(map[asset.Token]string)
	for _, token := range tokenWith200DollarPrincipal {
		token200Map[token] = ""
	}

	for i, pnl := range pnls {
		principalVal, ok := tokenWithDollarPrincipal[pnl.Token]
		if ok {
			pricipal, _ := new(big.Float).SetString(principalVal)
			pnls[i].PrincipalValueInDollar = pricipal
			continue
		}
		_, ok = token200Map[pnl.Token]
		if ok {
			pricipal, _ := new(big.Float).SetString("200")
			pnls[i].PrincipalValueInDollar = pricipal
			continue
		}

		fity, _ := new(big.Float).SetString("10")
		if pnl.FreeValue.Cmp(fity) == -1 {
			pricipal, _ := new(big.Float).SetString("1")
			pnls[i].PrincipalValueInDollar = pricipal
			continue
		}

		pricipal, _ := new(big.Float).SetString("100")
		pnls[i].PrincipalValueInDollar = pricipal
	}
}
