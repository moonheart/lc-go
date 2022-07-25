package main

import "strings"

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l <= r; {
		var bl byte = 0
		for {
			if bl1, ok := checkChar(s[l]); !ok {
				l++
				if l > r {
					break
				}
			} else {
				bl = bl1
				break
			}
		}

		var br byte = 0
		for {
			if br1, ok := checkChar(s[r]); !ok {
				r--
				if l > r {
					break
				}
			} else {
				br = br1
				break
			}

		}

		if bl != br {
			return false
		}
		l++
		r--
	}
	return true
}

func checkChar(byte2 byte) (byte, bool) {
	if byte2 < 48 {
		return 0, false
	}
	if byte2 < 58 {
		return byte2, true
	}
	if byte2 < 65 {
		return 0, false
	}
	if byte2 < 91 {
		return byte2, true
	}
	if byte2 < 97 {
		return 0, false
	}
	if byte2 < 123 {
		return byte2 - 32, true
	}
	return 0, false
}

func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	for l, r := 0, len(s)-1; l < r; {
		for l < r && !isChar(s[l]) {
			l++
		}
		for l < r && !isChar(s[r]) {
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isChar(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z')
}
