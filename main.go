package main

import "AlgorithmGolang/leetCode/week2"

func main() {
	//week1.IsPalindrome(121)
	//q := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//week2.MaxArea(q)

	//Lc T15
	//q := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//week2.ThreeSum(q)

	//LcT14
	//strs := []string{"ab", "a"}
	//week2.LongestCommonPrefix(strs)

	//LcT18
	q := []int{1, 0, -1, 0, -2, 2}
	n := 0
	week2.FourSum(q, n)
}
