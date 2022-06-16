package main

import (
	"fmt"
	"reflect"
)

func generate(numRows int) (res [][]int) {
	res = make([][]int, numRows)
	for i := range res {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return
}

func main() {
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
		if gotRes := generate(tt.args.numRows); !reflect.DeepEqual(gotRes, tt.wantRes) {
			fmt.Errorf("generate() = %v, want %v", gotRes, tt.wantRes)
		}
	}
}
