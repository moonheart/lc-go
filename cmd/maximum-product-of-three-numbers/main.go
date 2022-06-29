package main

import (
	"math"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	// three positive
	product := math.MinInt
	if nums[l-3] > 0 {
		product = nums[l-1] * nums[l-2] * nums[l-3]
	}
	// one positive, two negative
	if nums[1] < 0 && nums[l-1] > 0 {
		p := nums[0] * nums[1] * nums[l-1]
		if p > product {
			product = p
		}
	}
	if product != math.MinInt {
		return product
	}
	// tree negative, with min abs
	// one negative, two positive, with min abs
	abs := []int{}
	for _, n := range nums {
		if n < 0 {
			abs = append(abs, -n)
		} else {
			abs = append(abs, n)
		}
	}
	sort.Ints(abs)
	if abs[0] == 0 {
		return 0
	}
	return abs[0] * abs[1] * abs[2] * -1
}

func maximumProduct1(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	a := nums[0] * nums[1] * nums[l-1]
	b := nums[l-1] * nums[l-2] * nums[l-3]
	if a > b {
		return a
	}
	return b
}

func maximumProduct2(nums []int) int {
	max1, max2, max3 := math.MinInt, math.MinInt, math.MinInt
	min1, min2 := math.MaxInt, math.MaxInt
	for _, n := range nums {
		if n > max1 {
			max1, max2, max3 = n, max1, max2
		} else if n > max2 {
			max2, max3 = n, max2
		} else if n > max3 {
			max3 = n
		}

		if n < min1 {
			min1, min2 = n, min1
		} else if n < min2 {
			min2 = n
		}
	}
	a := min1 * min2 * max1
	b := max1 * max2 * max3
	if a > b {
		return a
	}
	return b
}

func main() {
	maximumProduct2([]int{-1, -2, -3})
}
