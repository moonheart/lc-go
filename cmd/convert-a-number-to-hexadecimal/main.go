package convert_a_number_to_hexadecimal

import "strings"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	sb := strings.Builder{}

	for i := 7; i >= 0; i-- {
		n := (num >> (4 * i)) & 0xf
		if n > 0 || sb.Len() > 0 {
			var digit byte
			if n >= 10 {
				digit = 'a' + byte(n-10)
			} else {
				digit = '0' + byte(n)
			}
			sb.WriteByte(digit)
		}
	}
	return sb.String()
}
