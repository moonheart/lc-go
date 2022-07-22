package main

import "sort"

func breakfastNumber(staple []int, drinks []int, x int) int {
	res := 0
	sort.Ints(staple)
	sort.Ints(drinks)
	idx := len(drinks) - 1
	for i := 0; i < len(staple); i++ {
		target := x - staple[i]
		for idx >= 0 && drinks[idx] > target {
			idx--
		}
		res += idx + 1
	}
	return res % 1000000007
}
