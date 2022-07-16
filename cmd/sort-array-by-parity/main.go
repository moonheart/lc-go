package main

func sortArrayByParity(nums []int) []int {
	l, r := 0, 0
	for l < len(nums) && r < len(nums) {
		if l == r && nums[l]&1 == 0 {
			l++
		} else if l != r && nums[r]&1 == 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
	return nums
}

func sortArrayByParity2(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l]&1 == 0 {
			l++
		}
		for l < r && nums[r]&1 == 1 {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return nums
}
