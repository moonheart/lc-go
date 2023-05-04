package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	arr := []int{}
	arr = pri(head, arr)
	return arr
}

func pri(node *ListNode, arr []int) []int {
	if node == nil {
		return arr
	}

	if node.Next != nil {
		arr = pri(node.Next, arr)
	}
	return append(arr, node.Val)
}
