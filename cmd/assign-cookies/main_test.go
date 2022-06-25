package main

import "testing"

func Test_findContentChildren(t *testing.T) {
	type args struct {
		g []int
		s []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"1", args{[]int{1, 2, 3}, []int{1, 1}}, 1},
		{"1", args{[]int{1, 2}, []int{1, 2, 3}}, 2},
		{"1", args{[]int{10, 9, 8, 7}, []int{5, 6, 7, 8}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findContentChildren(tt.args.g, tt.args.s); gotRes != tt.wantRes {
				t.Errorf("findContentChildren() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
