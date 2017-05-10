package utilx

import (
	"sort"
	"strconv"
	"strings"
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

// 对切片进行排序
func SliceSort() {}
func SliceSortString(s []string) []string {
	size := len(s)
	intSlice := make([]int, size)
	for i := 0; i < size; i++ {
		num, _ := strconv.Atoi(s[i])
		intSlice[i] = num
	}
	sort.Ints(intSlice)
	for i := 0; i < size; i++ {
		num := strconv.Itoa(intSlice[i])
		s[i] = num
	}

	return s
}

// 对切片排序，并保持索引关系
func SliceAsortDesc(s []int) []int {
	n := len(s)
	if n < 1 {
		return nil
	}

	index := make([]int, n)
	for i := 0; i < n; i++ {
		index[i] = i
	}

	n = n - 1
	for i := 0; i < n; i++ {
		m := n - i
		for j := 0; j < m; j++ {
			if s[j] < s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				index[j], index[j+1] = index[j+1], index[j]
			}
		}
	}
	return index
}

func SliceAsortDescInt64(s []int64) []int {
	n := len(s)
	if n < 1 {
		return nil
	}

	index := make([]int, n)
	for i := 0; i < n; i++ {
		index[i] = i
	}

	n = n - 1
	for i := 0; i < n; i++ {
		m := n - i
		for j := 0; j < m; j++ {
			if s[j] < s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				index[j], index[j+1] = index[j+1], index[j]
			}
		}
	}
	return index
}

func SliceAsortASCInt64(s []int64) []int {
	n := len(s)
	if n < 1 {
		return nil
	}

	index := make([]int, n)
	for i := 0; i < n; i++ {
		index[i] = i
	}

	n = n - 1
	for i := 0; i < n; i++ {
		m := n - i
		for j := 0; j < m; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				index[j], index[j+1] = index[j+1], index[j]
			}
		}
	}
	return index
}

// 常用于redis的hgetall查询出的结果
func SliceToMap(s []string) map[string]string {
	n := len(s)
	m := make(map[string]string)
	for i := 0; i < n; i += 2 {
		key := s[i]
		val := s[i+1]
		m[key] = val
	}
	return m
}

func StringTointSlice(list string) []int {
	s := strings.Split(list, ",")
	n := len(s)
	rs := make([]int, n)
	for i := 0; i < n; i++ {
		rs[i], _ = strconv.Atoi(s[i])
	}
	return rs
}
