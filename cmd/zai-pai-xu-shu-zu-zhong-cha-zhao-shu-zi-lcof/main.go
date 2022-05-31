package main

func search(nums []int, target int) int {
	var n int
	for i := len(nums) / 2; i < len(nums); i++ {
		if target == nums[i] {
			n++
		}
	}
	return n
}
