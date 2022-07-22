package main

import "testing"

func Test_firstPalindrome(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]string{"abc", "car", "ada", "racecar", "cool"}}, "ada"},
		{"2", args{[]string{"notapalindrome", "racecar"}}, "racecar"},
		{"3", args{[]string{"def", "ghi"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstPalindrome(tt.args.words); got != tt.want {
				t.Errorf("firstPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
