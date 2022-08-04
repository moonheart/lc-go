package backspace_string_compare

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"a##c", "#a#c"}, true},
		{"1", args{"ab#c", "ad#c"}, true},
		{"1", args{"ab##", "c#d#"}, true},
		{"1", args{"a#c", "b"}, false},
		{"1", args{"bxj##tw", "bxo#j##tw"}, true},
		{"1", args{"bxj##tw", "bxj###tw"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
