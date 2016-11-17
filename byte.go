package utilx

import (
	"strconv"
	// "strings"
)

func ByteToString(data []byte) string {
	return string(data)
}

func ByteToInt(data []byte) int {
	val, _ := strconv.Atoi(string(data))
	return val
}

func ByteToInt64(data []byte) int64 {
	val, _ := strconv.ParseInt(string(data), 10, 64)
	return val
}