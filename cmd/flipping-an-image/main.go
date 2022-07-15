package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		l, r := 0, len(row)-1
		for l < r {
			if row[l] == row[r] {
				row[l] ^= 1
				row[r] ^= 1
			}
			l++
			r--
		}
		if l == r {
			row[l] ^= 1
		}
	}
	return image
}

func flipAndInvertImage1(image [][]int) [][]int {
	println("start")
	print(image)
	l := len(image)
	for i := 0; i <= l/2; i++ {
		x := i
		for x >= 0 {
			if x < l/2 {
				print(image)
				fmt.Printf("%v,%v <> %v,%v\n", i, x, i, l-x-1)
				image[i][x], image[i][l-x-1] = revert(image[i][l-x-1]), revert(image[i][x])
				print(image)
				fmt.Printf("%v,%v <> %v,%v\n", l-i-1, x, l-i-1, l-x-1)
				image[l-i-1][x], image[l-i-1][l-x-1] = revert(image[l-i-1][l-x-1]), revert(image[l-i-1][x])
			}
			x--
		}
	}
	return image
}
func revert(a int) int {
	if a == 0 {
		return 1
	}
	return 0
}

func print(image [][]int) {
	for _, ints := range image {
		fmt.Println(ints)
	}
}
