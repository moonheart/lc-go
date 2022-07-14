package main

func shortestToChar(s string, c byte) (ans []int) {
	l, r1, r2 := 0, 0, 0
	for l < len(s) {
		if r2 >= l {
			if s[r2] == c {
				if s[r1] != c {
					ans = append(ans, r2-l)
				} else {
					ans = append(ans, min(r2-l, l-r1))
				}
				l++
			} else if r2 == len(s)-1 {
				ans = append(ans, l-r1)
				l++
			} else {
				r2++
			}
		} else {
			r1 = r2
			r2++
		}
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
