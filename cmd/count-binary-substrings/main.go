package count_binary_substrings

func countBinarySubstrings(s string) int {
	start := 0
	cur := 0
	res := 0
	for cur < len(s) && s[cur] == s[start] {
		cur++
	}
	mid := cur
	for cur < len(s) {
		for cur < len(s) && s[cur] != s[start] {
			cur++
		}
		res += min(cur-mid, mid-start)
		start = mid
		mid = cur
	}
	return res
}
func countBinarySubstrings1(s string) int {
	var cur, prev, sum int
	for cur < len(s) {
		start := s[cur]
		count := 0
		for cur < len(s) && s[cur] == start {
			cur++
			count++
		}
		sum += min(prev, count)
		prev = count
	}
	return sum
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
