package main

const mod int = 1e9 + 7

func numWays(n int) int {
	a, b, r := 1, 1, 0
	for i := 0; i < n; i++ {
		r = (a + b) % mod
		a = b
		b = r
	}
	return a
}

func main() {
	numWays(3)
}
