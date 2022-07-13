package main

func pivotIndex(nums []int) int {
	l := len(nums)
	f, b := make([]int, l), make([]int, l)

	for i := 0; i < l; i++ {
		if i == 0 {
			f[i] = 0
		} else {
			f[i] = f[i-1] + nums[i-1]
		}
	}
	for i := l - 1; i >= 0; i-- {
		if i == l-1 {
			b[i] = 0
		} else {
			b[i] = b[i+1] + nums[i+1]
		}
	}

	for i := 0; i < l; i++ {
		if f[i] == b[i] {
			return i
		}
	}
	return -1
}
func pivotIndex2(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for i, num := range nums {
		if sum*2+num == total {
			return i
		}
		sum += num
	}
	return -1
}
