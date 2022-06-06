package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) (res string) {
	if root == nil {
		return
	}
	res += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return
	}
	res += "("
	if root.Left != nil {
		res += tree2str(root.Left)
	}
	res += ")"
	if root.Right != nil {
		res += "("
		res += tree2str(root.Right)
		res += ")"
	}
	return
}

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 3},
	}
	println(tree2str(node))
}
