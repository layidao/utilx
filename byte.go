package utilx

import (
	"strconv"
	// "strings"
)

func ByteToString(data []byte) string {
	return string(data)
}

func ByteToInt(data []byte) int {
	val, err := strconv.Atoi(string(data))
	if err != nil {
		return 0
	}
	return val
}

func ByteToInt64(data []byte) int64 {
	val, _ := strconv.ParseInt(string(data), 10, 64)
	return val
}
