package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		if arr[l] != arr[r] {
			return false
		}
	}
	return true
}
