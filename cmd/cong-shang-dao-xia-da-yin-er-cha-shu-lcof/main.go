package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var arr []int
	if root == nil {
		return arr
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		node := nodes[0]
		arr = append(arr, node.Val)
		nodes = nodes[1:]
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
	return arr
}
