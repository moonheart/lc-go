package main

import (
	"reflect"
	"testing"
)

func Test_shortestToChar(t *testing.T) {
	type args struct {
		s string
		c byte
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{"1", args{"loveleetcode", 'e'}, []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
		{"1", args{"loveleetcode", 'l'}, []int{0, 1, 2, 1, 0, 1, 2, 3, 4, 5, 6, 7}},
		{"1", args{"aaab", 'b'}, []int{3, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := shortestToChar(tt.args.s, tt.args.c); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("shortestToChar() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
