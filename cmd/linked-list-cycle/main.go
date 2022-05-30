package main

type ListNode struct {
	Val  int
	Next *ListNode
}

var m = map[*ListNode]int{}

func hasCycle(head *ListNode) bool {
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = 0
		head = head.Next
	}
	return false
}
