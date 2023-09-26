package dateutil

import (
	"fmt"
	"strconv"
	"time"
)

func GetMonthStartDay() string {
	return time.Now().Format("2006-01") + "-01 00:00:00"
}

func GetMonthEndDay() string {
	month, _ := strconv.Atoi(time.Now().Format("01"))
	return time.Now().Format("2006-01") + "-" + fmt.Sprintf("%02d", GetYearMonthDay(time.Now().Year(), month)) + " 23:59:59"
}

func MothDay() []string {
	var res []string
	month, _ := strconv.Atoi(time.Now().Format("01"))
	day := GetYearMonthDay(time.Now().Year(), month)
	for i := 1; i <= day; i++ {
		res = append(res, strconv.Itoa(i))
	}
	return res
}

// GetYearMonthDay 获取当前月有多少天
func GetYearMonthDay(year, month int) int {
	day := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			// 闰年
			day = 29
		} else {
			// 平年
			day = 28
		}
	}
	return day
}

func DateFormat(date time.Time) (format string) {
	if !date.IsZero() {
		format = date.Format("2006-01-02 15:04:05")
	}
	return
}
