package main

import "sort"

func sortedSquares1(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}
	sort.Ints(nums)
	return nums
}
func sortedSquares(nums []int) (res []int) {
	res = make([]int, len(nums))
	l, r := 0, len(nums)-1
	i := len(nums) - 1
	for l <= r {
		if abs(nums[l]) > abs(nums[r]) {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
		i--
	}
	return res
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
