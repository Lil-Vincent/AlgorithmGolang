package week1

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	t := x
	res := 0
	for t != 0 {
		res = res*10 + t%10
		t /= 10
	}
	return x == res
}
