package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LeveldTreeNode struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var tmp []int
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmp)
	}

	return res
}
