package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	rootVal := root.Val
	second := -1

	var search func(node *TreeNode)
	search = func(node *TreeNode) {
		if node == nil || second != -1 && node.Val > second {
			return
		}
		if node.Val > rootVal {
			second = node.Val
		}
		search(node.Left)
		search(node.Right)
	}
	search(root)
	return second
}
