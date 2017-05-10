package utilx

import (
	"fmt"
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

func ByteToFloat32(data []byte) float32 {
	val, _ := strconv.ParseFloat(string(data), 32)
	fmt.Printf("============ %T, %v, %T, %v \n", val, val, float32(val), float32(val))
	return float32(val)
}

func ByteToFloat64(data []byte) float64 {
	val, _ := strconv.ParseFloat(string(data), 64)
	return val
}
