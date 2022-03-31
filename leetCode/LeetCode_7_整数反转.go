package main

import "math"

func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x /= 10
		if y < math.MinInt32 || y > math.MaxInt32 {
			return 0
		}
	}
	return y
}
