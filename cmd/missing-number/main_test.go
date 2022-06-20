package main

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 0, 1}}, 2},
		{"1", args{[]int{0, 1}}, 2},
		{"1", args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
		{"1", args{[]int{0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_missingNumber2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 0, 1}}, 2},
		{"1", args{[]int{0, 1}}, 2},
		{"1", args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
		{"1", args{[]int{0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber2(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_missingNumber3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantXor int
	}{
		{"1", args{[]int{3, 0, 1}}, 2},
		{"1", args{[]int{0, 1}}, 2},
		{"1", args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
		{"1", args{[]int{0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotXor := missingNumber3(tt.args.nums); gotXor != tt.wantXor {
				t.Errorf("missingNumber3() = %v, want %v", gotXor, tt.wantXor)
			}
		})
	}
}
