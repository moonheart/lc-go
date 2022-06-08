package main

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	k int
	sort.IntSlice
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		k: k,
	}
	for _, v := range nums {
		kl.Add(v)
	}
	return kl
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	k := 3
	nums := []int{5, -1}
	obj := Constructor(k, nums)
	println(obj.Add(2))
	println(obj.Add(1))
	println(obj.Add(-1))
	println(obj.Add(3))
	println(obj.Add(4))
}
