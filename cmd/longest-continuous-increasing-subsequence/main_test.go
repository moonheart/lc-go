package main

import "testing"

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 5, 4, 7}}, 3},
		{"1", args{[]int{1, 3, 5, 6, 7}}, 5},
		{"2", args{[]int{2, 2, 2, 2, 2}}, 1},
		{"3", args{[]int{2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
