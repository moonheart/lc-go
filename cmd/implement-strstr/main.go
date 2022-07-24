package main

func strStr(haystack string, needle string) int {
	i, j := 0, 0
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			if j == len(needle)-1 {
				return i
			}
			i1, j1 := i+1, j+1
			for i1 < len(haystack) && j1 < len(needle) && haystack[i1] == needle[j1] {
				if j1 == len(needle)-1 {
					return i
				}
				i1++
				j1++
			}
		}
		i++
	}
	return -1
}
func strStr2(haystack string, needle string) int {
outer:
	for i := 0; i+len(needle) <= len(haystack); i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}
