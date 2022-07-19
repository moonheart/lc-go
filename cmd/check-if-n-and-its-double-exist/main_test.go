package main

import "testing"

func Test_checkIfExist(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{10, 2, 5, 3}}, true},
		{"2", args{[]int{7, 1, 14, 11}}, true},
		{"3", args{[]int{3, 1, 7, 11}}, false},
		{"3", args{[]int{-2, 0, 10, -19, 4, 6, -8}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfExist2(tt.args.arr); got != tt.want {
				t.Errorf("checkIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
