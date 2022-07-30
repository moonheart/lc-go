package reverse_string_ii

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"abcdefg", args{"abcdefg", 2}, "bacdfeg"},
		{"abcd", args{"abcd", 2}, "bacd"},
		{"abcd/3", args{"abcd", 3}, "cbad"},
		{"a/2", args{"a", 2}, "a"},
		{"abcdefg/8", args{"abcdefg", 8}, "gfedcba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
