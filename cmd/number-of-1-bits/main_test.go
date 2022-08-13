package number_of_1_bits

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		n uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{0b00000000000000000000000000001011}, 3},
		{"1", args{0b00000000000000000000000010000000}, 1},
		{"1", args{0b11111111111111111111111111111101}, 31},
		{"1", args{85723}, 11},
		{"1", args{256}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.n); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
