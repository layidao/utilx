package utilx

import (
	"crypto/md5"
	// "crypto/sha1"
	"fmt"
	// "io"
	// "io/ioutil"
	"math/rand"
	// "net"
	// "net/http"
	// "net/url"
	"strconv"
	"strings"
	"time"
)

// 获取当前时间的字符串
func NewNonceString() string {
	nonce := strconv.FormatInt(time.Now().UnixNano(), 36)
	return fmt.Sprintf("%x", md5.Sum([]byte(nonce)))
}

// 获取定长随机数字符串
func GetWidthRandNum(base int) string {
	format := fmt.Sprintf("%%0%dd", len(strconv.Itoa(base)))
	return fmt.Sprintf(format, GetRandNum(base))
}

// 获取token
func GetToken() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	rndNum := rand.Int63()
	uuid := Md5(Md5(strconv.FormatInt(nano, 10)) + Md5(strconv.FormatInt(rndNum, 10)))
	return uuid
}

// 字符串长度
func Strlen(str string) int {
	return len([]rune(str))
}

// IP 地址转为正数
func IP2Int(ip string) int64 {
	bits := strings.Split(ip, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}
