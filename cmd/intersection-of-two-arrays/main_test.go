package main

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name  string
		args  args
		wantA []int
	}{
		{"1", args{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2}},
		{"1", args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, []int{4, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(gotA, tt.wantA) {
				t.Errorf("intersection() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}
