package number_complement

func findComplement(num int) int {
	// 持续寻找最高位
	/*
		0000000010101001
		        ^
	*/
	highBit := 0
	for i := 1; i <= 30; i++ {
		if num < 1<<i {
			break
		}
		highBit = i
	}
	/*
		0000000010101001
		0000000010000000 = 1<<(highBit)
		0000000100000000 = 1<<(highBit+1)
		0000000011111111 = 1<<(highBit+1)-1 = mask

		0000000010101001 ^
		0000000011111111 = 0000000001010110
	*/
	mask := 1<<(highBit+1) - 1
	return num ^ mask
}
