package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	idx := (len(nums) - 1) / 2
	mid := nums[idx]
	if target == mid {
		return idx
	} else if idx == 0 && len(nums) == 1 {
		return -1
	} else if target > mid {
		i := search(nums[idx+1:], target)
		if i == -1 {
			return -1
		}
		return idx + i + 1
	} else {
		return search(nums[:idx], target)
	}
}

func search1(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		n := nums[mid]
		if target == n {
			return mid
		}
		if target > n {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"4", args{[]int{-1, 0, 3, 5, 9, 12}, 0}, 1},
	}
	for _, tt := range tests {
		if got := search(tt.args.nums, tt.args.target); got != tt.want {
			fmt.Errorf("search() = %v, want %v", got, tt.want)
		}
	}
}
