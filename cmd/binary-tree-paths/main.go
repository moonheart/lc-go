package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var arr []string
	dfs(root, "", &arr)
	return arr
}

func dfs(node *TreeNode, path string, arr *[]string) {
	if len(path) == 0 {
		path = fmt.Sprintf("%d", node.Val)
	} else {
		path = fmt.Sprintf("%s->%d", path, node.Val)
	}
	if node.Left == nil && node.Right == nil {
		*arr = append(*arr, path)
		return
	}
	if node.Left != nil {
		dfs(node.Left, path, arr)
	}
	if node.Right != nil {
		dfs(node.Right, path, arr)
	}
}

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	paths := binaryTreePaths(node)
	fmt.Print(paths)
}
