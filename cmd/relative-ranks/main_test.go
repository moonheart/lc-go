package main

import (
	"reflect"
	"testing"
)

func Test_findRelativeRanks(t *testing.T) {
	type args struct {
		score []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"1", args{[]int{5, 4, 3, 2, 1}}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
		{"1", args{[]int{10, 3, 8, 9, 4}}, []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findRelativeRanks(tt.args.score); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("findRelativeRanks() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
