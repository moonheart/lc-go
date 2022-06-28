package main

func maxCount(m int, n int, ops [][]int) int {
	mx, my := m, n
	for _, op := range ops {
		if op[0] < mx {
			mx = op[0]
		}
		if op[1] < my {
			my = op[1]
		}
	}
	return mx * my
}
