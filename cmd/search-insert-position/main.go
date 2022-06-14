package main

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := ((r - l) / 2) + l
		if target > nums[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

func main() {
	nums := []int{1, 3}
	println(searchInsert(nums, 2))
}
