package reverse_string_ii

func reverseStr(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(s); i += k * 2 {
		sub := b[i:min(len(s), i+k)]
		for j, n := 0, len(sub); j < n/2; j++ {
			sub[j], sub[n-1-j] = sub[n-1-j], sub[j]
		}
	}
	return string(b)
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
