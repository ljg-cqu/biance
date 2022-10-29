package utils

import (
	"fmt"
	"github.com/ljg-cqu/biance/utils/sign"
)

func CalculateAndAppendSignature(msg, key string) string {
	mac := sign.CalculateMAC(msg, key)
	msg += fmt.Sprintf("&signature=%v", mac)
	return msg
}
