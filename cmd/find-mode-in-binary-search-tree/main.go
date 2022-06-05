package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var answer []int
var pre = 0
var count = 0
var max = 0

func findMode(root *TreeNode) []int {
	dfs(root)
	answer = nil
	pre = 0
	count = 0
	max = 0
	return answer
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Left)
	if node.Val == pre {
		count++
	} else {
		pre, count = node.Val, 1
	}
	if count == max {
		answer = append(answer, node.Val)
	} else if count > max {
		max = count
		answer = []int{node.Val}
	}
	dfs(node.Right)
}

func main() {
	node := &TreeNode{Val: 0}
	fmt.Print(findMode(node))
}
