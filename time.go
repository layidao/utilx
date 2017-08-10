package utilx

import (
	"fmt"
	"strconv"
	"time"
)

const ChinaTimeZoneOffset = 8 * 60 * 60 //Beijing(UTC+8:00)
// 时间戳的字符串
func TimestampString() string {
	return fmt.Sprintf("%d", time.Now().Unix()+ChinaTimeZoneOffset)
}

// 获取当前时间戳
func GetCurrentTime() int {
	return int(time.Now().Unix())
}

// 获取当前日期
func GetCurrentDate(format string) string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	if format == "" {
		format = "2006-01-02"
	}
	return tm.Format(format)
}

// 今天的开始时间错
func GetCurrentDayStartTime() int64 {
	currDate := GetCurrentDate("")
	currDate = currDate + " 00:00:00"
	return Date2Stamp(currDate, "2006-01-02 15:04:05")
}

// 时间戳转日期格式
func StampStringToDate(timestampStr, format string) string {
	timestamp, _ := strconv.ParseInt(timestampStr, 10, 64)
	tm := time.Unix(timestamp, 0)

	if format == "" {
		format = "2006-01-02 15:04:05"
	}

	return tm.Format(format)
}

// 时间戳转日期
func StampIntToDate(timestamp int64, format string) string {
	tm := time.Unix(timestamp, 0)

	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return tm.Format(format)
}

// 日期转时间戳
func Date2Stamp(date string, format string) int64 {
	// tm, _ := time.Parse("2006-01-02 15:04:05", "2016-05-08 00:00:00")
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	tm, _ := time.Parse(format, date)
	return tm.Unix()
}

// 获取前n天的数据
func NDateBefore(n int, isMulti bool, format string) []string {
	rs := make([]string, 0)
	nTime := time.Now()

	if format == "" {
		format = "20060102"
	}

	if isMulti {
		j := -1
		for i := 0; i < n; i++ {
			yesTime := nTime.AddDate(0, 0, j)
			j--

			logDay := yesTime.Format(format)
			rs = append(rs, logDay)
		}
	} else {
		yesTime := nTime.AddDate(0, 0, -1*n)
		logDay := yesTime.Format(format)
		rs = append(rs, logDay)
	}
	return rs
}
