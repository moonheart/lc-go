package reverse_only_letters

func reverseOnlyLetters(s string) string {
	b := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !(s[l] >= 'A' && s[l] <= 'Z' || s[l] >= 'a' && s[l] <= 'z') {
			l++
		}
		for l < r && !(s[r] >= 'A' && s[r] <= 'Z' || s[r] >= 'a' && s[r] <= 'z') {
			r--
		}
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}
