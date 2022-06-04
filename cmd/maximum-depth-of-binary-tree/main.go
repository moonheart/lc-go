package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lmax := maxDepth(root.Left)
	rmax := maxDepth(root.Right)
	if lmax > rmax {
		return lmax + 1
	} else {
		return rmax + 1
	}
}
