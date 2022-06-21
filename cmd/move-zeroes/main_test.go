package main

import "testing"

func Benchmark_moveZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{0, 0, 0, 1, 2, 3, 0, 0, 0, 1}
		moveZeroes(nums)
	}
}

func Benchmark_moveZeroes1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{0, 0, 0, 1, 2, 3, 0, 0, 0, 1}
		moveZeroes1(nums)
	}
}
