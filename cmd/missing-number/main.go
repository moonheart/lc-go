package main

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i, n := range nums {
		if i != n {
			return i
		}
	}
	return len(nums)
}

func missingNumber2(nums []int) int {
	m := map[int]bool{}
	for _, n := range nums {
		m[n] = true
	}
	for i := 0; ; i++ {
		if !m[i] {
			return i
		}
	}
}

func missingNumber3(nums []int) (xor int) {
	for i, n := range nums {
		xor ^= i ^ n
	}
	return xor ^ len(nums)
}
