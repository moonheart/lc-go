package main

func nextGreaterElement(nums1 []int, nums2 []int) (res []int) {
	m := map[int]int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		if i == len(nums2)-1 {
			m[nums2[i]] = -1
		} else {
			for j := i + 1; j < len(nums2); j++ {
				if nums2[j] > nums2[i] {
					m[nums2[i]] = nums2[j]
					break
				}
			}
			if _, ok := m[nums2[i]]; !ok {
				m[nums2[i]] = -1
			}
		}
	}
	for _, i := range nums1 {
		res = append(res, m[i])
	}
	return
}

func nextGreaterElement2(nums1 []int, nums2 []int) (res []int) {
	m := map[int]int{}
	var stack []int
	for i := len(nums2) - 1; i >= 0; i-- {
		n := nums2[i]
		for len(stack) > 0 && n >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			m[n] = -1
		} else {
			m[n] = stack[len(stack)-1]
		}
		stack = append(stack, n)
	}
	for _, i := range nums1 {
		res = append(res, m[i])
	}
	return
}
