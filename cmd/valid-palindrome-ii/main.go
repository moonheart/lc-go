package valid_palindrome_ii

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] == s[r] {
			l++
			r--
			continue
		}
		rl, rr := true, true
		for i, j := l, r-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				rl = false
				break
			}
		}
		for i, j := l+1, r; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				rr = false
				break
			}
		}
		return rl || rr
	}
	return true
}

func validPalindromex(s string) bool {
	return validPalindrome1([]byte(s), false)
}

func validPalindrome1(s []byte, removed bool) bool {
	//fmt.Printf("%v, %v\n", string(s), removed)
	if len(s) <= 1 {
		return true
	}
	if s[0] == s[len(s)-1] {
		return validPalindrome1(s[1:len(s)-1], removed)
	}
	if removed {
		return false
	}
	return validPalindrome1(s[0:len(s)-1], true) || validPalindrome1(s[1:len(s)], true)
}
