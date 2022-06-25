package main

import "math"

func minMoves(nums []int) int {
	sum := 0
	min := math.MaxInt
	for _, num := range nums {
		sum += num
		if num < min {
			min = num
		}
	}
	return sum - len(nums)*min
}
