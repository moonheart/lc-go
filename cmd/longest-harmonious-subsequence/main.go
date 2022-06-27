package main

import "fmt"

func findLHS(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	max := 0
	for i, _ := range m {
		if m[i+1] > 0 {
			c := m[i] + m[i+1]
			if c > max {
				max = c
			}
		}
	}
	return max
}

func findLHS1(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}
	for num, c := range cnt {
		fmt.Printf("cur: %d, +1: %d\n", num, num+1)
		if c1 := cnt[num+1]; c1 > 0 && c+c1 > ans {
			ans = c + c1
		}
	}
	return
}

func main() {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 2, 2, 5, 2, 3, 7, 6}}, 5},
		//{"1", args{[]int{1, 2, 3, 4}}, 2},
		//{"1", args{[]int{1, 1, 1, 1}}, 0},
	}
	for _, tt := range tests {
		if got := findLHS1(tt.args.nums); got != tt.want {
			fmt.Errorf("findLHS() = %v, want %v", got, tt.want)
		}
	}
}
