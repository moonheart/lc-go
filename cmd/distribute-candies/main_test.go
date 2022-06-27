package main

import "testing"

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candyType []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"1", args{[]int{1, 1, 2, 2, 3, 3}}, 3},
		{"2", args{[]int{1, 1, 2, 3}}, 2},
		{"3", args{[]int{6, 6, 6, 6}}, 1},
		{"4", args{[]int{1000, 1000, 2, 1, 2, 5, 3, 1}}, 4},
		{"4", args{[]int{0, 0, 0, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := distributeCandies(tt.args.candyType); gotRes != tt.wantRes {
				t.Errorf("distributeCandies() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
