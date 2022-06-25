package main

import "testing"

func Test_islandPerimeter(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"1", args{[][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
		}}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := islandPerimeter2(tt.args.grid); gotRes != tt.wantRes {
				t.Errorf("islandPerimeter() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		islandPerimeter([][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
		})
	}
}
func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		islandPerimeter2([][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
		})
	}
}
