package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, n := range nums {
		if v, ok := m[n]; ok && i-v <= k {
			return true
		} else {
			m[n] = i
		}
	}
	return false
}

func containsNearbyDuplicate1(nums []int, k int) bool {
	m := map[int]struct{}{}
	for i, n := range nums {
		if i > k {
			delete(m, nums[i-k-1])
		}
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = struct{}{}
	}
	return false
}
