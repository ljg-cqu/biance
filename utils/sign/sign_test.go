package sign

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	msg    = "symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559"
	msgMAC = "c8db56825ae71d6d79447849e617115f4a920fa2acdcab2b053c4b2838bd6b71"
	key    = "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j"
)

func TestCalculateMAC(t *testing.T) {
	mac := CalculateMAC(msg, key)
	require.Equal(t, msgMAC, mac)
}

func TestValidMAC(t *testing.T) {
	msg := "symbol=LTCBTC&side=BUY&type=LIMIT&timeInForce=GTC&quantity=1&price=0.1&recvWindow=5000&timestamp=1499827319559"
	msgMAC := "c8db56825ae71d6d79447849e617115f4a920fa2acdcab2b053c4b2838bd6b71"
	key := "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j"
	isValid := ValidateMAC(msg, msgMAC, key)
	require.True(t, isValid)
}
