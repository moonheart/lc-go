package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	maxAvg := float64(math.MinInt)
	for i := 0; i < len(nums); i++ {
		if i < k {
			sum += nums[i]
		} else {
			sum += nums[i] - nums[i-k]
		}
		if i >= k-1 {
			avg := float64(sum) / float64(k)
			if avg > maxAvg {
				maxAvg = avg
			}
		}
	}
	return maxAvg
}

func findMaxAverage1(nums []int, k int) float64 {
	sum := 0
	for _, n := range nums[:k] {
		sum += n
	}
	maxSum := sum
	for i, n := range nums[k:] {
		sum += n - nums[i]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}

func main() {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{[]int{1, 12, -5, -6, 50, 3}, 4}, 12.75},
		{"1", args{[]int{5}, 1}, 5.0},
	}
	for _, tt := range tests {
		if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
			fmt.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
		}
	}
}
