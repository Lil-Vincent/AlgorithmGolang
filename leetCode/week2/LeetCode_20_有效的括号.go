package week2

import (
	"math"
)

func IsValid(s string) bool {
	tt := 0
	n := len(s)
	stk := make([]byte, n+1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			tt++
			stk[tt] = s[i]
		} else {
			if tt != 0 && math.Abs(float64(s[i]-stk[tt])) <= 2 {
				tt--
			} else {
				return false
			}
		}
	}
	return tt == 0
}
