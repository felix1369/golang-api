package utils

import (
	"strconv"
	"time"
)

func Chunks(xs []string, chunkSize int) [][]string {
	if len(xs) == 0 {
		return nil
	}
	divided := make([][]string, (len(xs)+chunkSize-1)/chunkSize)
	prev := 0
	i := 0
	till := len(xs) - chunkSize
	for prev < till {
		next := prev + chunkSize
		divided[i] = xs[prev:next]
		prev = next
		i++
	}
	divided[i] = xs[prev:]
	return divided
}

// StrToFloat64 convert string to float64
func StrToFloat64(str string) (value *float64, err error) {
	*value, err = strconv.ParseFloat(str, 64)
	return value, err
}

// StrToInt64 convert string to int64
func StrToInt64(str string) (value *int64, err error) {
	*value, err = strconv.ParseInt(str, 10, 64)
	return value, err
}

// StrToDate convert string to LocalTim
func StrToDate(str string) (value *LocalTime, err error) {
	err = value.ScanString(str)
	return value, err
}

// StrToFloat64Value convert string to float64
func StrToFloat64Value(str string) *float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return &value
	}
	return nil
}

// StrToInt64Value convert string to int64
func StrToInt64Value(str string) *int64 {
	value, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return &value
	}
	return nil
}

// StrToUIntValue convert string to int64
func StrToUIntValue(str string) *uint {
	value, err := strconv.ParseUint(str, 10, 32)
	if err == nil {
		uintVal := uint(value)
		return &uintVal
	}
	return nil
}

// StrToDateValue convert string to LocalTim
func StrToDateValue(str string) *LocalTime {
	value, err := time.Parse("01/02/2006", str)
	if err == nil {
		return &LocalTime{Time: value}
	}
	return nil
}

// StrToDateTimeValue2 convert string to LocalTim
func StrToDateTimeValue2(strDateTime string) *LocalTime {
	value, err := time.Parse("01/02/2006 15:04:05 PM", strDateTime)
	if err == nil {
		return &LocalTime{Time: value}
	}
	return nil
}

// StrToDateTimeValue convert string to LocalTim
func StrToDateTimeValue(strDate string, strTime string) *LocalTime {
	strDateTime := strDate + " " + strTime
	value, err := time.Parse("01/02/2006 15:04:05 PM", strDateTime)
	if err == nil {
		return &LocalTime{Time: value}
	}
	return nil
}

// StrToDateTimeValue convert string to LocalTim
func StrToTimeValue(strDate string, strTime string) *time.Time {
	strDateTime := strDate + " " + strTime
	value, err := time.Parse("01/02/2006 15:04:05 PM", strDateTime)
	if err == nil {
		return &value
	}
	return nil
}
