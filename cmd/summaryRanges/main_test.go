package main

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"1", args{[]int{0, 1, 2, 4, 5, 7}}, []string{"0->2", "4->5", "7"}},
		{"1", args{[]int{0, 2, 3, 4, 6, 8, 9}}, []string{"0", "2->4", "6", "8->9"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := summaryRanges(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("summaryRanges() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
