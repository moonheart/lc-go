package main

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1}, 1}, 0},
		{"2", args{[]int{1}, 0}, 0},
		{"3", args{[]int{1}, 2}, 1},
		{"4", args{[]int{1, 3}, 2}, 1},
		{"5", args{[]int{1, 3}, 0}, 0},
		{"6", args{[]int{1, 3}, 1}, 0},
		{"7", args{[]int{1, 3}, 3}, 1},
		{"8", args{[]int{1, 3}, 4}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
