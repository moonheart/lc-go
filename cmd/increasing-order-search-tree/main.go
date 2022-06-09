package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var inorder func(node *TreeNode)
	res := &TreeNode{}
	cur := res
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		cur.Right = &TreeNode{Val: node.Val}
		cur = cur.Right
		inorder(node.Right)
	}
	inorder(root)
	return res.Right
}

func main() {
	node := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 7},
	}
	fmt.Print(increasingBST(node))
}
