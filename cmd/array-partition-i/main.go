package main

import "sort"

func arrayPairSum(nums []int) (res int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return
}
