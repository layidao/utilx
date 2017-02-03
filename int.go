package utilx

import (
	"math/rand"
	"time"
)

func GetRandNum(base int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(base)
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else if x == 0 {
		return 0
	}
	return x
}
