package tools

import "strconv"

func FormatFloat(i float64) string {
	return strconv.FormatFloat(i, 'f', -1, 64)
}

func ParseFloat(s string) float64 {
	i, _ := strconv.ParseFloat(s, 64)
	return i
}
