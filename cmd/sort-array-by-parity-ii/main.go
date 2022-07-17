package main

func sortArrayByParityII(nums []int) []int {
	res := make([]int, len(nums))
	a, b := 0, 1
	for _, num := range nums {
		if num&1 == 0 {
			res[a] = num
			a += 2
		} else {
			res[b] = num
			b += 2
		}
	}
	return res
}
func sortArrayByParityII2(nums []int) []int {
	for i, j := 0, 1; i < len(nums); i += 2 {
		if nums[i]&1 == 1 {
			for nums[j]&1 == 1 {
				j += 2
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
