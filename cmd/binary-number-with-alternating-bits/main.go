package binary_number_with_alternating_bits

func hasAlternatingBits(n int) bool {
	/*
		10101010 = n
		01010101 = n>>1
		10101010 ^
		01010101 =
		11111111 = a
	*/
	a := n ^ n>>1
	/*
		 11111111 = a
		100000000 = a + 1
		000000000 = a&(a+1)
	*/
	return a&(a+1) == 0
}
