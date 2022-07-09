package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{-1, 0, 3, 5, 9, 12}, 9}, 4},
		{"2", args{[]int{-1, 0, 3, 5, 9, 12}, 2}, -1},
		{"3", args{[]int{-1, 0, 3, 5, 9}, 2}, -1},
		{"4", args{[]int{-1, 0, 3, 5, 9, 12}, 0}, 1},
		{"5", args{[]int{-1, 0, 3, 5, 9, 12}, -1}, 0},
		{"6", args{[]int{-1, 0, 3, 5, 9, 12}, 12}, 5},
		{"7", args{[]int{-1, 0}, 12}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
