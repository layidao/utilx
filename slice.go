package utilx

import (
	"strconv"
)

func SliceContainsInt(sl []int, v int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func SliceContainsString(sl []string, v string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// 获取slice中最大的数值，此方法只适用于所有元素为数字型元素
func SliceMaxString(s []string) int {
	size := len(s)
	if size == 0 {
		return 0
	}

	var max int
	for i := 0; i < size; i++ {
		num, _ := strconv.Atoi(s[i])
		if num > max {
			max = num
		}
	}

	return max
}

// 取切片中的数据唯一
func SliceUniqueString(s []string) []string {
	size := len(s)
	if size == 0 {
		return []string{}
	}

	m := make(map[string]bool)
	for i := 0; i < size; i++ {
		m[s[i]] = true
	}

	realLen := len(m)
	ret := make([]string, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}

	return ret
}
