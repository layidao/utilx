package utilx

import (
	"math/rand"
	"time"
)

func GetRandNum(base int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(base)
}
