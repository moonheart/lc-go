package main

var chars = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func reverseVowels(s string) string {
	bytes := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		for l < len(s) && !chars[s[l]] {
			l++
		}
		for r > -1 && !chars[s[r]] {
			r--
		}
		if l < r {
			bytes[l], bytes[r] = bytes[r], bytes[l]
		}
	}
	return string(bytes)
}
