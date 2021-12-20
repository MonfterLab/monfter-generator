package utils

import (
	"reflect"
	"sort"
	"time"
)

const Location = "Asia/Shanghai"
const OriginalDate = "2006-01-02"
const OriginalDateTime = "2006-01-02 15:04:05"
const OriginalExcelTime = "2006/1/2 15:04" //Excel date time format
const OriginalTimeStr = "20060102150405"
const OriginalMonthDay = "01月02日" //n月n日

const TimeoutLimit = 10

func Now() int {
	return int(time.Now().Unix())
}

func NowUTC() int {
	return int(time.Now().UTC().Unix())
}

// GetNowDateTime
func GetNowDateTime() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

func StrToTime(datetime string) int {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, time.Local)
	return int(t.Unix())
}

func StrToUTCTime(datetime string) int {
	t, _ := time.Parse("2006-01-02 15:04:05", datetime)
	return int(t.Unix())
}

func GetNowDate() (date string) {
	date = time.Now().UTC().Format("20060102")
	return
}

// FindByStringArray the array must be ordered array
func FindByStringArray(value string, array []string) int {
	index := -1

	if len(array) == 0 {
		return index
	}

	sort.Strings(array)

	v := reflect.ValueOf(value).String()

	left := 0
	right := len(array) - 1
	for {
		mid := (left + right) / 2
		if reflect.ValueOf(array[mid]).String() == v {
			index = mid
			break
		}
		if left == right {
			break
		}
		if reflect.ValueOf(array[left]).String() <= reflect.ValueOf(array[mid]).String() {
			if reflect.ValueOf(array[left]).String() <= v && v < reflect.ValueOf(array[mid]).String() {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if reflect.ValueOf(array[mid]).String() <= v && v <= reflect.ValueOf(array[right]).String() {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return index
}

// FindByIntArray the array must be ordered array
func FindByIntArray(value int, array []int) int {
	index := -1

	if len(array) == 0 {
		return index
	}
	sort.Ints(array)

	left := 0
	right := len(array) - 1
	for {
		mid := (left + right) / 2
		if array[mid] == value {
			index = mid
			break
		}
		if left == right {
			break
		}
		if array[left] <= array[mid] {
			if array[left] <= value && value < array[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if array[mid] <= value && value <= array[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return index
}
