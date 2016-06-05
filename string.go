package utilx

import (
	"crypto/md5"
	// "crypto/sha1"
	"fmt"
	// "io"
	// "io/ioutil"
	// "math/rand"
	// "net"
	// "net/http"
	// "net/url"
	"strconv"
	// "strings"
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

// 字符串长度
func Strlen(str string) int {
	return len([]rune(str))
}
