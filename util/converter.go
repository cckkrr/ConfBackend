package util

import "strconv"

func Float64ToInt64(f float64) int64 {
	return int64(f)
}

func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func Int64ToFloat64(i int64) float64 {
	return float64(i)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
