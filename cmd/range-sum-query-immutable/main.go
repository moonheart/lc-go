package main

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}
	return NumArray{sums}
}

func (a *NumArray) SumRange(left int, right int) int {
	return a.sums[right+1] - a.sums[left]
}
