package time

import (
	"strconv"
	"time"
)

func Timestamp() string {
	return strconv.Itoa(int(time.Now().UnixMilli()))
}
