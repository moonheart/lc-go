package counting_bits

const (
	m1  = 0x55555555 // 01010101010101010101010101010101
	m2  = 0x33333333 // 00110011001100110011001100110011
	m4  = 0x0f0f0f0f // 00001111000011110000111100001111
	m8  = 0x00ff00ff // 00000000111111110000000011111111
	m16 = 0x0000ffff // 00000000000000001111111111111111
)

func countBits1(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		x := (i>>1)&m1 + i&m1
		//fmt.Printf("%32b %d %b\n", i, i, x)
		x = (x>>2)&m2 + x&m2
		//fmt.Printf("%32b %d %b\n", i, i, x)
		x = (x>>4)&m4 + x&m4
		//fmt.Printf("%32b %d %b\n", i, i, x)
		x = (x>>8)&m8 + x&m8
		//fmt.Printf("%32b %d %b\n", i, i, x)
		x = (x>>16)&m16 + x&m16
		//fmt.Printf("%32b %d %b\n", i, i, x)
		//x = (x * 0x01010101) >> 24
		//fmt.Printf("%32b %d %b\n", i, i, x)
		//if x > 32 {
		//panic("ssdsd")
		//}
		res[i] = x
	}
	return res
}

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i>>1] + i&1
		// bits[i] = bits[i&(i-1)] + 1
	}
	return res
}
