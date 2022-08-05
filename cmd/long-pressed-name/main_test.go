package long_pressed_name

import "testing"

func Test_isLongPressedName(t *testing.T) {
	type args struct {
		name  string
		typed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"alex", "aaleex"}, true},
		{"1", args{"saeed", "ssaaedd"}, false},
		{"1", args{"vtkgn", "vttkgnn"}, true},
		{"1", args{"rick", "kric"}, false},
		{"1", args{"alex", "aaleexa"}, false},
		{"1", args{"leelee", "lleeelee"}, true},
		{"1", args{"alexd", "ale"}, false},
		{"1", args{"pyplrz", "ppyypllr"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLongPressedName(tt.args.name, tt.args.typed); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
