package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func CalculateMAC(msg, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func ValidateMAC(msg, msgMAC, key string) bool {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(msg))
	expectedMAC := mac.Sum(nil)
	msgMACByt, _ := hex.DecodeString(msgMAC)
	return hmac.Equal(msgMACByt, expectedMAC)
}

func calculateMAC(msg, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	return mac.Sum(nil)
}

func validateMAC(msg, msgMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(msgMAC, expectedMAC)
}
