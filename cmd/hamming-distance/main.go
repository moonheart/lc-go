package hamming_distance

import "math/bits"

func hammingDistance1(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func hammingDistance(x int, y int) (res int) {
	for i := x ^ y; i > 0; i &= i - 1 {
		res++
	}
	return
}
