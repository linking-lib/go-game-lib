package strs

import (
	"fmt"
	"strconv"
)

func IsEmpty(str string) bool {
	return len(str) == 0
}

func IntToStr(value int) string {
	return strconv.Itoa(value)
}

func StrToInit(value string) int {
	v, _ := strconv.Atoi(value)
	return v
}

func Int64ToStr(value int64) string {
	return strconv.FormatInt(value, 10)
}

func StrToInt64(value string) int64 {
	v, _ := strconv.ParseInt(value, 10, 64)
	return v
}

func FloatToStr(value float64) string {
	return fmt.Sprintf("%.6f", value)
}

func FloatToStrByScale(value float64, scale string) string {
	return fmt.Sprintf("%."+scale+"f", value)
}

func StrToFloat(value string) float64 {
	v, _ := strconv.ParseFloat(value, 64)
	return v
}

func BoolToStr(value bool) string {
	return strconv.FormatBool(value)
}

func StrToBool(value string) bool {
	v, _ := strconv.ParseBool(value)
	return v
}

func InterfaceToStr(value interface{}) string {
	return value.(string)
}

func ByteToStr(value interface{}) string {
	return string(value.([]byte))
}
