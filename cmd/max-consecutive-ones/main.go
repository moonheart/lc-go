package main

func findMaxConsecutiveOnes(nums []int) int {
	x := 0
	max := 0
	for _, num := range nums {
		if num == 1 {
			x++
			if x > max {
				max = x
			}
		} else {
			x = 0
		}
	}
	return max
}

func findMaxConsecutiveOnes2(nums []int) int {
	max := 0
	for i := range nums {
		j := i
		for j < len(nums) && nums[j] == 1 {
			j++
		}
		if j-i > max {
			max = j - i
		}
		i = j
	}
	return max
}
