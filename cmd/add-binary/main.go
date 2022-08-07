package add_binary

import (
	"strconv"
)

func addBinary(a string, b string) string {
	prev := 0
	res := ""
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 || prev > 0 {
		sum := 0
		if i >= 0 && a[i] == '1' {
			sum += 1
		}
		if j >= 0 && b[j] == '1' {
			sum += 1
		}
		sum += prev
		if sum >= 2 {
			prev = 1
		} else {
			prev = 0
		}
		res = strconv.Itoa(sum&1) + res
		i--
		j--
	}
	return res
}
