package main

import (
	"sort"
)

func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[i]+nums[j] > target {
			j--
		}
		res += j - i
		i++
	}
	return res % 1000000007
}
