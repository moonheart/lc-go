package everse_words_in_a_string_iii

func reverseWords(s string) string {
	b := []byte(s)
	for i, j := 0, 0; i < len(s) && j < len(s); j++ {
		if b[j] == ' ' {
			reverse(b[i:j])
			i = j + 1
		} else if j == len(s)-1 {
			reverse(b[i : j+1])
			i = j + 1
		}
	}
	return string(b)
}

func reverseWords1(s string) string {
	b := []byte(s)
	i := 0
	for i < len(s) {
		l := i
		for i < len(s) && s[i] != ' ' {
			i++
		}
		r := i - 1
		for l < r {
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
		for i < len(s) && s[i] == ' ' {
			i++
		}
	}
	return string(b)
}

func reverse(arr []byte) {
	for j, n := 0, len(arr); j < n/2; j++ {
		arr[j], arr[n-1-j] = arr[n-1-j], arr[j]
	}
}
