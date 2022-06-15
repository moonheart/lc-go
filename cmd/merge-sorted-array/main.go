package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, i2, j := m-1, m+n-1, n-1
	for i >= 0 || j >= 0 {
		if j >= 0 && i >= 0 {
			if nums2[j] >= nums1[i] {
				nums1[i2] = nums2[j]
				j--
			} else {
				nums1[i2] = nums1[i]
				i--
			}
		} else if j >= 0 {
			nums1[i2] = nums2[j]
			j--
		} else {
			nums1[i2] = nums1[i]
			i--
		}
		i2--
	}
}

func main() {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3}, []int{1, 2, 2, 3, 5, 6}},
		//{"2", args{nums1: []int{1}, m: 1, nums2: []int{}, n: 0}, []int{1}},
		//{"3", args{nums1: []int{0}, m: 0, nums2: []int{1}, n: 1}, []int{1}},
	}
	for _, tt := range tests {
		merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		if !testEq1(tt.args.nums1, tt.want) {
			fmt.Errorf("merge() results %v, want %v", tt.args.nums1, tt.want)
		}
	}
}

func testEq1(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
