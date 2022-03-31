package main

import (
	"math"
)

func myAtoi(s string) int {
	k := 0
	for k < len(s) && s[k] == ' ' {
		k++
	}
	if k == len(s) {
		return 0
	}
	minus := 1
	if s[k] == '-' {
		minus = -1
		k++
	} else if s[k] == '+' {
		k++
	}

	res := 0
	for k < len(s) && s[k] >= '0' && s[k] <= '9' {
		x := s[k] - '0'
		if minus > 0 && res > (math.MaxInt32-int(x))/10 {
			return math.MaxInt32
		}
		if minus < 0 && -res < (math.MinInt32+int(x))/10 {
			return math.MinInt32
		}
		if minus < 0 && res*10+int(x) == math.MaxInt32+1 {
			return math.MinInt32
		}
		res = res*10 + int(s[k]) - int('0')
		k++

	}

	res *= minus
	if res > 0 && res > math.MaxInt32 {
		return math.MaxInt32
	}
	if res < 0 && res < math.MinInt32 {
		return math.MinInt32 * minus
	}
	return res
}
