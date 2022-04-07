package week2

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)

	var res [][]int
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l+1 < r && nums[l] == nums[l+1] {
					l++
					continue
				}
				for r-1 < l && nums[r] == nums[r-1] {
					r--
					continue
				}
				l++
				r--
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
