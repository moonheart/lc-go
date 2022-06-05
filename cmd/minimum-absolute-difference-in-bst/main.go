package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	min, pre := math.MaxInt, -1
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node.Left != nil {
			inOrder(node.Left)
		}
		if pre != -1 && node.Val-pre < min {
			min = node.Val - pre
		}

		pre = node.Val

		if node.Right != nil {
			inOrder(node.Right)
		}
	}

	inOrder(root)
	return min
}

func main() {
	node := &TreeNode{
		Val: 236,
		Left: &TreeNode{
			Val:   104,
			Right: &TreeNode{Val: 227},
		},
		Right: &TreeNode{
			Val:   701,
			Right: &TreeNode{Val: 911},
		},
	}
	println(getMinimumDifference(node))
}
