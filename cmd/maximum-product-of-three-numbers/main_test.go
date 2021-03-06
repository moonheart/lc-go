package main

import "testing"

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3}}, 6},
		{"2", args{[]int{1, 2, 3, 4}}, 24},
		{"3", args{[]int{-1, -2, -3}}, -6},
		{"4", args{[]int{-1, -2, -3, 0, 1, 2, 3}}, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct2(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
