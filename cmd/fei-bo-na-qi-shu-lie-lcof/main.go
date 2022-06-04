package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	const mod int = 1e9 + 7
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p, q = q, r
		r = (p + q) % mod
	}
	return r
}

func main() {
	fmt.Print(fib(43))
}
