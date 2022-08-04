package backspace_string_compare

func backspaceCompare(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		i = gonext(s, i)
		j = gonext(t, j)

		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}

func gonext(s string, i int) int {
	skip := 0
	for i >= 0 {
		if s[i] == '#' {
			skip++
			i--
		} else if skip > 0 {
			skip--
			i--
		} else {
			break
		}
	}
	return i
}
