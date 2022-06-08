package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	var search func(node *TreeNode)
	min, pre := math.MaxInt, -1
	search = func(node *TreeNode) {
		if node == nil {
			return
		}
		search(node.Left)
		if pre != -1 && node.Val-pre < min {
			min = node.Val - pre
		}
		pre = node.Val
		search(node.Right)
	}
	search(root)
	return min
}

func main() {
	node := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 6},
	}
	println(minDiffInBST(node))
}
