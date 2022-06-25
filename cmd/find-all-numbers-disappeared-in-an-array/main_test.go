package main

import (
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findDisappearedNumbers2(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	}
}
func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDisappearedNumbers2([]int{4, 3, 2, 7, 8, 2, 3, 1})
	}
}
