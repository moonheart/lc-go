package main

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParity(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 1, 2, 4}}, []int{2, 4, 3, 1}},
		{"2", args{[]int{0, 1, 2}}, []int{0, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParity(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArrayByParity([]int{23, 23, 43, 534, 2345, 6, 2745, 248, 345, 234, 2345, 34, 457, 23, 34, 345, 2345, 7, 845, 457, 245, 234, 356, 43, 234, 27, 23})
	}
}
func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortArrayByParity2([]int{23, 23, 43, 534, 2345, 6, 2745, 248, 345, 234, 2345, 34, 457, 23, 34, 345, 2345, 7, 845, 457, 245, 234, 356, 43, 234, 27, 23})
	}
}
