package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	sum := 0
	for i := 0; i < len(flowerbed) && sum < n; {
		if flowerbed[i] == 0 {
			if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
				sum++
				i += 2
			} else {
				i += 1
			}
		} else {
			i += 2
		}

	}
	return sum >= n
}

func main() {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 0, 0, 0, 0, 1}, 2}, false},
		{"1", args{[]int{1, 0, 0, 0, 1, 0, 0}, 2}, true},
		{"1", args{[]int{1, 0, 0, 0, 1}, 1}, true},
		{"2", args{[]int{1, 0, 0, 0, 1}, 2}, false},
	}
	for _, tt := range tests {
		if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
			fmt.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
		}
	}
}
