package main

func findErrorNums(nums []int) []int {
	arr := make([]int, len(nums))
	res := []int{}
	for _, num := range nums {
		if arr[num-1] != 0 {
			res = append(res, num)
		}
		arr[num-1] = 1
	}
	for i, i2 := range arr {
		if i2 == 0 {
			res = append(res, i+1)
		}
	}
	return res
}
