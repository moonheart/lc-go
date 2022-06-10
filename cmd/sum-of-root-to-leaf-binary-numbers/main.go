package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	var dfs func(node *TreeNode, val int)
	var sum int
	dfs = func(node *TreeNode, val int) {
		if node.Left == nil && node.Right == nil {
			sum += val + node.Val
			return
		}
		if node.Left != nil {
			dfs(node.Left, (val+node.Val)<<1)
		}
		if node.Right != nil {
			dfs(node.Right, (val+node.Val)<<1)
		}
	}
	dfs(root, 0)
	return sum
}
