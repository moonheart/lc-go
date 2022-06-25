package main

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 1, 0, 1, 1, 1}}, 3},
		{"2", args{[]int{1, 0, 1, 1, 0, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})
	}
}
func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMaxConsecutiveOnes2([]int{1, 1, 0, 1, 1, 1})
	}
}

/*
go test -bench . -benchmem
goos: windows
goarch: amd64
cpu: AMD Ryzen 5 5600X 6-Core Processor
Benchmark1-12           442477549                2.512 ns/op           0 B/op          0 allocs/op
Benchmark2-12           179057521                6.653 ns/op           0 B/op          0 allocs/op
PASS
ok      _/D_/Git_Github/lcof-go/cmd/max-consecutive-ones        3.292s
*/
