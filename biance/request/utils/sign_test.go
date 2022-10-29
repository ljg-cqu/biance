package utils

import (
	"fmt"
	time2 "github.com/ljg-cqu/biance/utils/time"
	"testing"
)

const (
	key = "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j"
)

func TestCalculateAndAppendSignature(t *testing.T) {
	body := fmt.Sprintf("asset=%v&timestamp=%v", "AVAX", time2.Timestamp())
	body = CalculateAndAppendSignature(body, key)
	fmt.Println(body)
}
