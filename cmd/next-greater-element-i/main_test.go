package main

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"1", args{[]int{4, 1, 2}, []int{1, 3, 4, 2}}, []int{-1, 3, -1}},
		{"1", args{[]int{2, 4}, []int{1, 2, 3, 4}}, []int{3, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := nextGreaterElement(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("nextGreaterElement() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
		nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4})
	}
}
func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nextGreaterElement2([]int{4, 1, 2}, []int{1, 3, 4, 2})
		nextGreaterElement2([]int{2, 4}, []int{1, 2, 3, 4})
	}
}
