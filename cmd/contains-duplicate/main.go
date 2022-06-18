package main

func containsDuplicate(nums []int) bool {
	m := map[int]struct{}{}
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = struct{}{}
	}
	return false
}
