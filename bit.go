package utilx

import (
	"strconv"
	"strings"
)

// 位运算相关,字典code转为64位的位置整数
func MoveBitInt(code int, bitLen int) (int, uint64) {
	if code == 0 {
		return 0, 0
	}

	code--
	group := code / bitLen
	bitIndex := code % bitLen

	var bitVal uint64
	bitVal = 1 << uint64(bitIndex)
	//log.Println("++++++", bitVal)
	return group, bitVal
}

//将 1,2,3转为 位运算字符串
func MoveBitString(val string, bitLen int) string {
	if val == "0" || val == "" {
		return "0"
	}

	cells := strings.Split(val, ",")
	cells = SliceUniqueString(cells)

	//log.Println(cells)

	maxElem := SliceMaxString(cells)
	rsLen := (maxElem / bitLen)
	if maxElem%bitLen > 0 {
		rsLen = rsLen + 1
	}

	rs := make([]uint64, rsLen)
	slen := len(cells)
	for i := 0; i < slen; i++ {
		code, _ := strconv.Atoi(cells[i])
		group, val := MoveBitInt(code, 60)
		rs[group] += val
	}

	// log.Println(rs)

	rsStrSlice := make([]string, rsLen)
	for i := 0; i < rsLen; i++ {
		// codeStr := strconv.Itoa(rs[i])
		codeStr := strconv.FormatUint(rs[i], 10)
		rsStrSlice[i] = codeStr
	}

	return strings.Join(rsStrSlice, "|")
}
