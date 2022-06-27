package week3

func NextPermutation(nums []int) {
	n := len(nums) - 1
	k := n
	for k > 0 && nums[k-1] >= nums[k] {
		k -= 1
	}
	if k == 0 {
		reverse(nums, 0, n)
	} else {
		l := k - 1
		r := k
		for r <= n && nums[r] > nums[l] {
			r += 1
		}
		nums[l], nums[r-1] = nums[r-1], nums[l]
		reverse(nums, l+1, n)
	}
}
func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
