package util

import (
	"fmt"
	"reflect"
	"strings"
)

// GetStructNames get struct variable name and combine into string
func GetStructNames(e reflect.Value) (names []string) {
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		names = append(names, varName)
	}
	return names
}

// GetStructValues get struct item value and combine into string
func GetStructValues(e reflect.Value) (values []interface{}) {
	for i := 0; i < e.NumField(); i++ {
		varValue := e.Field(i).Interface()
		values = append(values, varValue)
	}
	return values
}

// GetStructInfos get struct item value and name then combine into seperate string
func GetStructInfos(e reflect.Value) (names []string, values []interface{}) {
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		names = append(names, varName)

		varValue := e.Field(i).Interface()
		values = append(values, varValue)
	}
	return names, values
}

// GetStructNamesIndexStart get struct variable name from a certain index and combine into string
func GetStructNamesIndexStart(e reflect.Value, indexStart int) (names []string) {
	for i := indexStart; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		names = append(names, varName)
	}
	return names
}

// GetStructValuesIndexStart get struct item value from a certain index  and combine into string
func GetStructValuesIndexStart(e reflect.Value, indexStart int) (values []interface{}) {
	for i := indexStart; i < e.NumField(); i++ {
		varValue := e.Field(i).Interface()
		values = append(values, varValue)
	}
	return values
}

// GetStructInfosIndexStart get struct item value and name from a certain index then combine into seperate string
func GetStructInfosIndexStart(e reflect.Value, indexStart int) (names []string, values []interface{}) {
	for i := indexStart; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		names = append(names, varName)

		varValue := e.Field(i).Interface()
		values = append(values, varValue)
	}
	return names, values
}

// GetStructNumField get struct number of field
func GetStructNumField(e reflect.Value) int {
	return e.NumField()
}

// GetStructInsertPlaceHolder get struct placeholder for insert query
func GetStructInsertPlaceHolder(e reflect.Value, indexStart int) string {
	tempPlaceHolder := []string{}
	for i := indexStart; i < e.NumField(); i++ {
		tempPlaceHolder = append(tempPlaceHolder, "?")
	}
	return fmt.Sprintf("%v", strings.Join(tempPlaceHolder, ", "))
}

// GetInsertPlaceHolder get struct placeholder for insert query
func GetInsertPlaceHolder(index int) string {
	tempPlaceHolder := []string{}
	for i := 0; i < index; i++ {
		tempPlaceHolder = append(tempPlaceHolder, "?")
	}
	return fmt.Sprintf("%v", strings.Join(tempPlaceHolder, ", "))
}

// GetStructValuesField get struct item value from a certain index  and combine into string
func GetStructValuesField(e reflect.Value, fields []string) (values []interface{}) {
	for i := 0; i < len(fields); i++ {
		varValue := e.FieldByName(fields[i]).Interface()
		values = append(values, varValue)
	}
	return values
}

func find(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Substring get substring
func Substring(s string, start int, end int) string {
	startStrIdx := 0
	i := 0
	for j := range s {
		if i == start {
			startStrIdx = j
		}
		if i == end {
			return s[startStrIdx:j]
		}
		i++
	}
	return s[startStrIdx:]
}

// StrConvert get string from array of interface
func StrConvert(sep string, a ...interface{}) string {
	str := ""
	lengthA := len(a)
	if lengthA > 0 {
		dataObj := a[0].([]interface{})
		lenDataObj := len(dataObj)
		if lenDataObj > 0 {
			str = fmt.Sprintf("%v", dataObj[0])
			for index := 1; index < lengthA; index++ {
				str1 := fmt.Sprintf("%v", a[index])
				str += sep + str1
			}
		}
	}
	return str
}
