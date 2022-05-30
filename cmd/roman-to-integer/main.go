package main

var vals = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		v := s[i]
		if i == len(s)-1 {
			res += vals[v]
		} else if i < len(s)-1 {
			next := vals[s[i+1]]
			if vals[v] < next {
				res += next - vals[v]
				i = i + 1
			} else {
				res += vals[v]
			}
		}
	}
	return res
}

func main() {
	romanToInt("MCMXCIV")
}
