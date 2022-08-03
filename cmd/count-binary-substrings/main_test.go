package count_binary_substrings

import "testing"

func Test_countBinarySubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"00110011"}, 6},
		{"1", args{"10101"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBinarySubstrings1(tt.args.s); got != tt.want {
				t.Errorf("countBinarySubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
