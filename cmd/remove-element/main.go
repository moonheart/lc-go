package main

import "fmt"

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)
	for i < j {
		if nums[i] == val {
			nums[i] = nums[j-1]
			j--
		} else {
			i++
		}
	}
	return i
}

func main() {
	ints := []int{4, 5}
	println(removeElement(ints, 5))
	fmt.Print(ints)
}
