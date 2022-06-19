package main

import (
	"strconv"
)

func summaryRanges(nums []int) (res []string) {
	for i := 0; i < len(nums); {
		start := i
		for i++; i < len(nums) && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[start])
		if start < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		res = append(res, s)
	}
	return
}
