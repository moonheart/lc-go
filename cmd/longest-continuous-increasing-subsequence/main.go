package main

func findLengthOfLCIS(nums []int) int {
	maxL := 1
	L := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			L++
			if L > maxL {
				maxL = L
			}
		} else {
			L = 1
		}
	}
	return maxL
}
