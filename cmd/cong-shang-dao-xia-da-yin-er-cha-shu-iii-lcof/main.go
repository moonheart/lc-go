package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		var tmp []int
		tmpq := queue
		queue = nil
		for _, node := range tmpq {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if level%2 == 1 {
			for i, j := 0, len(tmp); i < j/2; i++ {
				tmp[i], tmp[j-i-1] = tmp[j-i-1], tmp[i]
			}
		}

		res = append(res, tmp)
	}

	return res
}

func main() {
	node := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 5},
		},
	}
	levelOrder(node)
}
