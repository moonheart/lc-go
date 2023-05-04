package main

import (
	"math"
	"testing"
)

func Test_strToInt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"123"}, 123},
		{"2", args{"12  3"}, 12},
		{"3", args{"-12  3"}, -12},
		{"3", args{"123-"}, 123},
		{"4", args{"- 12  3"}, 0},
		{"4", args{"134527349857353"}, math.MaxInt32},
		{"4", args{"-134527349857989353"}, math.MinInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToInt(tt.args.str); got != tt.want {
				t.Errorf("strToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
