package utilx

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

// 获取一个 0-maxInt 之间的随机整数, 不包含maxInt
func RandIntn(maxInt int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxInt)
}

// 非4舍5入，保留2位数的除法，division
func DivisionN(dividend, divisor float64, n int) float64 {
	if dividend == 0 || divisor == 0 {
		return float64(0)
	}

	m := math.Pow10(n)
	return float64(int64((dividend/divisor)*m)) / m
}

// 对float进行四舍五入
func Float64Round(x float64, n int) float64 {
	times := math.Pow10(n)
	x = math.Round(x * times)
	return Float2Float(x/times, n)
}

func Float64Floor(x float64, n int) float64 {
	times := math.Pow10(n)
	x = math.Floor(x * times)
	return Float2Float(x/times, n)
}

// 概率元素
func ProbabilityElems(elems []float64, precise float64) (band []int) {
	// 先排序找到最小的数
	sort.Slice(elems, func(i, j int) bool {
		return elems[i] < elems[j]
	})

	// 取最小的数的位数
	n := math.Ceil(math.Log10(precise / elems[0]))
	// 通过位数求倍数
	p := math.Pow10(int(n))

	l := len(elems)
	band = make([]int, l)
	band[0] = int(elems[0] * p)
	for i := 1; i < l; i++ {
		// 累加
		band[i] = band[i-1] + int(elems[i]*p)
	}
	return
}

// 获得一个概率索引
func Probability(band []int) (index, r int) {
	l := len(band)
	max := band[l-1]
	r = RandIntn(max)
	for i, v := range band {
		if r <= v {
			return i, r
		}
	}

	index = l - 1
	r = max
	return
}

// float精度
func Float2Float(num float64, precise int) float64 {
	format := fmt.Sprintf("%%.%df", precise)
	fnum, _ := strconv.ParseFloat(fmt.Sprintf(format, num), 64)
	return fnum
}

// 获取给定两个值的中间的随机数
func RangeRandFloat64(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Float64()
	return min + ((max - min) * r)
}
