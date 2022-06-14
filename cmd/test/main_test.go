package main

import "testing"

func BenchmarkFind1(b *testing.B) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		findNumberIn2DArray1(matrix, 1) // run fib(30) b.N times
	}
}

func BenchmarkFind2(b *testing.B) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		findNumberIn2DArray2(matrix, 1) // run fib(30) b.N times
	}
}
