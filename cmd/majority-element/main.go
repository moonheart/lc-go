package main

func majorityElement(nums []int) int {
	x := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			x = nums[i]
		}
		if nums[i] == x {
			count++
		} else {
			count--
		}
	}
	return x
}
