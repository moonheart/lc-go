package main

func missingNumber(nums []int) int {
	if nums[0] != 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			return nums[i] + 1
		}
		if nums[i]+1 != nums[i+1] {
			return nums[i] + 1
		}
	}
	return -1
}
