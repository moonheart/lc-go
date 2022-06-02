package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix)-1, 0
	for m >= 0 && n < len(matrix[0]) {
		if target == matrix[m][n] {
			return true
		}
		if target > matrix[m][n] {
			n++
		} else {
			m--
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	print(findNumberIn2DArray(matrix, 30))
}
