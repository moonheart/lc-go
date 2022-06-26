package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}
	rs := make([][]int, r)
	m, n := 0, 0
	for _, ints := range mat {
		for _, num := range ints {
			if n >= c {
				m++
				n = 0
			}
			rs[m] = append(rs[m], num)
			n++
		}
	}
	return rs
}
func matrixReshape1(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}
	rs := make([][]int, r)
	for i := range rs {
		rs[i] = make([]int, c)
	}
	for i := 0; i < r*c; i++ {
		rs[i/c][i%c] = mat[i/len(mat[0])][i%len(mat[0])]
	}
	return rs
}
