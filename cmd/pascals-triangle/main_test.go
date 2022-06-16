package main

import (
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		{"test0", args{numRows: 0}, [][]int{}},
		{"test1", args{numRows: 1}, [][]int{{1}}},
		{"test2", args{numRows: 2}, [][]int{{1}, {1, 1}}},
		{"test3", args{numRows: 3}, [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{"test4", args{numRows: 4}, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
		{"test5", args{numRows: 5}, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := generate(tt.args.numRows); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("generate() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
