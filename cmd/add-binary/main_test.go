package add_binary

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"11", "1"}, "100"},
		{"1", args{"1010", "1011"}, "10101"},
		{"1", args{"1111", "1111"}, "11110"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
