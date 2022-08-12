package power_of_four

var m1 = 0x55555555 // 01010101010101010101010101010101

func isPowerOfFour(n int) bool {
	return (n > 0) && ((n & (n - 1)) == 0) && (m1&n == n)
}
