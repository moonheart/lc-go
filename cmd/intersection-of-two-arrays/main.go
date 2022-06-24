package main

import "sort"

func intersection(nums1 []int, nums2 []int) (a []int) {
	m1 := map[int]struct{}{}
	m2 := map[int]struct{}{}
	for _, i := range nums1 {
		m1[i] = struct{}{}
	}
	for _, i := range nums2 {
		m2[i] = struct{}{}
	}

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	for i := range m1 {
		if _, ok := m2[i]; ok {
			a = append(a, i)
		}
	}
	return
}
func intersection2(nums1 []int, nums2 []int) (a []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if len(a) == 0 || x > a[len(a)-1] {
				a = append(a, x)
			}
			i++
			j++
		} else if x > y {
			j++
		} else {
			i++
		}
	}
	return
}
