package utils

import "strings"

func StrInArrayIndex(str string, array []string) int {
	st := strings.Trim(str, " ")
	for i, v := range array {
		if v == st {
			return i
		}
	}
	return -1
}

func IntInArrayIndex(x int, array []int) int {
	for i, v := range array {
		if v == x {
			return i
		}
	}
	return -1
}

func Int64InArrayIndex(x int64, array []int64) int {
	for i, v := range array {
		if v == x {
			return i
		}
	}
	return -1
}
