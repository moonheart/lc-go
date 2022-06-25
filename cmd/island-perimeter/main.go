package main

func islandPerimeter(grid [][]int) (res int) {
	for i, ints := range grid {
		for j, num := range ints {
			if num == 1 {
				if i == 0 || grid[i-1][j] == 0 {
					res++
				}
				if j == 0 || ints[j-1] == 0 {
					res++
				}
				if i == len(grid)-1 || grid[i+1][j] == 0 {
					res++
				}
				if j == len(ints)-1 || ints[j+1] == 0 {
					res++
				}
			}
		}
	}
	return
}

func dfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 1
	}
	if grid[i][j] == 0 {
		return 1
	}
	if grid[i][j] == 2 {
		return 0
	}
	grid[i][j] = 2
	return dfs(grid, i-1, j) +
		dfs(grid, i+1, j) +
		dfs(grid, i, j-1) +
		dfs(grid, i, j+1)
}
func islandPerimeter2(grid [][]int) (res int) {
	for i, ints := range grid {
		for j, num := range ints {
			if num == 1 {
				return dfs(grid, i, j)
			}
		}
	}
	return
}
