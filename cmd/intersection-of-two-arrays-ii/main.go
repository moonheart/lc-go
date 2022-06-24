package main

import "sort"

func intersect(nums1 []int, nums2 []int) (a []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			a = append(a, x)
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
