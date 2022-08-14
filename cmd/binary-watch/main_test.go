package binary_watch

import (
	"reflect"
	"testing"
)

func Test_readBinaryWatch(t *testing.T) {
	type args struct {
		turnedOn int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"1", args{1}, []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := readBinaryWatch(tt.args.turnedOn); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("readBinaryWatch() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
