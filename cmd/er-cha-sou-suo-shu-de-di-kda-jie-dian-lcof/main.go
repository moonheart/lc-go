package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	var dfs func(node *TreeNode)
	var res int
	dfs = func(node *TreeNode) {
		if node == nil || k == 0 {
			return
		}
		dfs(node.Right)
		k--
		if 0 == k {
			res = node.Val
		}
		dfs(node.Left)
	}
	dfs(root)
	return res
}

func main() {
	node := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 4},
	}
	println(kthLargest(node, 1))
}
