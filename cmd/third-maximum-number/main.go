package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)
	k := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			k++
		}
		if k == 3 {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

func thirdMax2(nums []int) int {
	var a, b, c *int
	for _, n := range nums {
		n := n
		if a == nil || n > *a {
			a, b, c = &n, a, b
		} else if *a > n && (b == nil || n > *b) {
			b, c = &n, b
		} else if b != nil && *b > n && (c == nil || n > *c) {
			c = &n
		}
	}
	if c != nil {
		return *c
	}
	return *a
}

func main() {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 2, 1}}, 1},
		{"1", args{[]int{1, 2}}, 2},
		{"1", args{[]int{2, 2, 3, 1}}, 1},
	}
	for _, tt := range tests {
		if got := thirdMax(tt.args.nums); got != tt.want {
			fmt.Errorf("thirdMax() = %v, want %v", got, tt.want)
		}
	}
}
