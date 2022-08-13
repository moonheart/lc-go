package find_the_difference

func findTheDifference(s string, t string) byte {
	m := make(map[byte]int, len(s))
	for _, i := range s {
		m[byte(i)]++
	}
	for _, i := range t {
		m[byte(i)]--
		if m[byte(i)] < 0 {
			return byte(i)
		}
	}
	return 0
}

func findTheDifference1(s string, t string) byte {
	var a int32
	for _, i := range s {
		a ^= i
	}
	for _, i := range t {
		a ^= i
	}
	return byte(a)
}
