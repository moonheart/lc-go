package main

import "testing"

func Test_calPoints(t *testing.T) {
	type args struct {
		ops []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"1", args{[]string{"5", "2", "C", "D", "+"}}, 30},
		{"1", args{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}}, 27},
		{"1", args{[]string{"1"}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := calPoints(tt.args.ops); gotRes != tt.wantRes {
				t.Errorf("calPoints() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
