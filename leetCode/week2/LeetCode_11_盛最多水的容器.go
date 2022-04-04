package week2

func MaxArea(height []int) int {

	res := 0
	for i, j := 0, len(height)-1; i < j; {
		area := min(height[i], height[j]) * (j - i)
		res = max(res, area)
		if height[i] >= height[j] {
			j--
		} else {
			i++
		}
	}
	return res
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
