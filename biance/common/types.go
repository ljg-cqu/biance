package common

import "math/big"

type Token string

type Symbol string

type Price struct {
	Symbol Symbol
	Price  *big.Float
}
