package main

import (
	"sort"
)

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

func breakfastNumber2(staple []int, drinks []int, x int) int {
	res := 0
	sort.Ints(staple)
	sort.Ints(drinks)
	i, j := 0, len(drinks)-1
	for i < len(staple) && j >= 0 {
		for j >= 0 && staple[i]+drinks[j] > x {
			j--
		}
		res += j + 1
		i++
	}
	return res % 1000000007
}
