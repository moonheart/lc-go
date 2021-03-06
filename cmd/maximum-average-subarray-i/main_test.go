package main

import "testing"

func Test_findMaxAverage(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverage1(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
