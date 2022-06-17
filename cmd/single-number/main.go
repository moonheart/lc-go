package main

func singleNumber(nums []int) int {
	x := 0
	for _, n := range nums {
		x ^= n
	}
	return x
}
