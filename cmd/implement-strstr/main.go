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
