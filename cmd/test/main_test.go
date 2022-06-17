package main

import "testing"

func BenchmarkFind1(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if n&1 == 0 {

		}
	}
}

func BenchmarkFind2(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if n%2 == 0 {
		}
	}
}
