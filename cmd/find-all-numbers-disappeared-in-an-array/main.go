package main

import (
	"fmt"
	"reflect"
)

func findDisappearedNumbers1(nums []int) (res []int) {
	x := make([]int, len(nums))
	for _, num := range nums {
		x[num-1] = num
	}
	for i, n := range x {
		if n == 0 {
			res = append(res, i+1)
		}
	}
	return
}

func findDisappearedNumbers2(nums []int) (res []int) {
	l := len(nums)
	for _, num := range nums {
		nums[(num-1)%l] += l
	}
	for i, n := range nums {
		if n <= l {
			res = append(res, i+1)
		}
	}
	return
}

func main() {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"1", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{5, 6}},
		{"1", args{[]int{1, 1}}, []int{2}},
	}
	for _, tt := range tests {
		if gotRes := findDisappearedNumbers2(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
			fmt.Errorf("findDisappearedNumbers() = %v, want %v", gotRes, tt.wantRes)
		}
	}
}
