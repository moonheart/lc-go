package remove_palindromic_subsequences

func removePalindromeSub(s string) int {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return 2
		}
	}
	return 1
}
