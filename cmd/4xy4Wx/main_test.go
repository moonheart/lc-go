package main

import "testing"

func Test_purchasePlans(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 5, 3, 5}, 6}, 1},
		{"1", args{[]int{2, 2, 1, 9}, 10}, 4},
		{"1", args{[]int{40330, 31957, 63879, 13204, 47923, 56282, 75126, 3423, 98483}, 60482}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := purchasePlans(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("purchasePlans() = %v, want %v", got, tt.want)
			}
		})
	}
}
