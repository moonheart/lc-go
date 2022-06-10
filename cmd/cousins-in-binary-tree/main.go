package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	queue := []*TreeNode{root}
	parents := map[*TreeNode]*TreeNode{}
	for len(queue) > 0 {
		queue1 := queue
		queue = nil
		var px *TreeNode
		var py *TreeNode
		for _, node := range queue1 {
			p, ok := parents[node]
			if node.Val == x && ok {
				px = p
			} else if node.Val == y && ok {
				py = p
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
				parents[node.Left] = node
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				parents[node.Right] = node
			}
		}

		if px != nil && py != nil && px != py {
			return true
		} else if px == nil && py != nil || px != nil && py == nil {
			return false
		}
	}
	return false
}
