package main

func imageSmoother(img [][]int) [][]int {
	res := make([][]int, len(img))
	for i := range res {
		res[i] = make([]int, len(img[0]))
	}
	m, n := len(img), len(img[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := img[i][j]
			c := 1
			if i > 0 {
				sum += img[i-1][j]
				c++
			}
			if i < m-1 {
				sum += img[i+1][j]
				c++
			}
			if j > 0 {
				sum += img[i][j-1]
				c++
			}
			if j < n-1 {
				sum += img[i][j+1]
				c++
			}
			if i > 0 && j > 0 {
				sum += img[i-1][j-1]
				c++
			}
			if i > 0 && j < n-1 {
				sum += img[i-1][j+1]
				c++
			}
			if i < m-1 && j > 0 {
				sum += img[i+1][j-1]
				c++
			}
			if i < m-1 && j < n-1 {
				sum += img[i+1][j+1]
				c++
			}
			res[i][j] = sum / c
		}
	}
	return res
}
