package main

import "fmt"

func moveZeroes(nums []int) {
	for z, i := 0, 1; i < len(nums); {
		if nums[z] == 0 && nums[i] != 0 {
			nums[z], nums[i] = nums[i], nums[z]
			z++
			i++
		} else if nums[z] == 0 && nums[i] == 0 {
			i++
		} else if nums[z] != 0 && nums[i] == 0 {
			z = i
			i++
		} else {
			z += 2
			i += 2
		}
	}
}

func moveZeroes1(nums []int) {
	for l, r := 0, 0; r < len(nums); {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}

func main() {
	//nums := []int{0, 1, 0, 3, 12}
	//moveZeroes(nums)
	//fmt.Println(nums)
	nums := []int{0, 0, 0, 1, 2, 3, 0, 0, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}
