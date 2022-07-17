package main

func diStringMatch(s string) []int {
	res := []int{}
	i, j := 0, len(s)
	for _, i3 := range s {
		if i3 == 'I' {
			res = append(res, i)
			i++
		} else {
			res = append(res, j)
			j--
		}
	}
	res = append(res, i)
	return res
}
