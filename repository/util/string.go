package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type StringUtil struct{}

var String StringUtil = StringUtil{}

func (_ StringUtil) IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

func (_ StringUtil) IsNull(str string) bool {
	return strings.ToLower(str) == "null"
}

func (_ StringUtil) HaveValueArray(strArray []string) bool {
	return strArray != nil || len(strArray) > 0
}

func GetExcelColumnName(columnNumber int) string {
	dividend := columnNumber
	var columnName string
	var modulo int

	for dividend > 0 {
		modulo = (dividend - 1) % 26
		columnName = toCharStr(modulo) + columnName
		dividend = int((dividend - modulo) / 26)
	}

	return columnName
}

func toCharStr(i int) string {
	return string(rune('A' + i))
}

func GetSQLStringValue(colVal interface{}) string {
	var result string
	switch v := colVal.(type) {
	case string:
		result = fmt.Sprintf("'%v'", v)
	case uint:
		result = fmt.Sprintf("%v", v)
	case *uint:
		if v == nil {
			result = "NULL"
		} else {
			result = fmt.Sprintf("%v", *v)
		}
	case int64:
		result = fmt.Sprintf("%v", v)
	case *int64:
		if v == nil {
			result = "NULL"
		} else {
			result = fmt.Sprintf("%v", *v)
		}
	case float64:
		result = fmt.Sprintf("%f", v)
	case *float64:
		if v == nil {
			result = "NULL"
		} else {
			result = fmt.Sprintf("%f", *v)
		}
	case LocalTime:
		result = fmt.Sprintf("'%v'", v.Format(time.RFC3339))
	case *LocalTime:
		if v == nil {
			result = "NULL"
		} else {
			local := *v
			result = fmt.Sprintf("'%v'", local.Format(time.RFC3339))
		}
	case time.Time:
		result = fmt.Sprintf("'%v'", v.Format(time.RFC3339))
	case *time.Time:
		if v == nil {
			result = "NULL"
		} else {
			local := *v
			result = fmt.Sprintf("'%v'", local.Format(time.RFC3339))
		}
	case uuid.UUID:
		result = fmt.Sprintf("'%v'::uuid", v)
	default:
		formater := "'%v'"
		val := fmt.Sprintf("%v", v)
		if val == "<nil>" {
			formater = "%v"
			val = "NULL"
		}
		result = fmt.Sprintf(formater, val)
	}
	return result
}

func GetCpOtherColumnName(columnName string) string {
	if len(columnName) > 0 {
		substring := columnName[:2]
		if strings.ToLower(substring) == "cp" {
			return columnName[2:]
		} else {
			return fmt.Sprintf("Cp%v", columnName)
		}
	}
	return columnName
}
