package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, x int) bool {
	var m = map[int]struct{}{}
	var find func(node *TreeNode, k int) bool
	find = func(node *TreeNode, k int) bool {
		if node == nil {
			return false
		}
		if _, ok := m[k-node.Val]; ok {
			return true
		}
		m[node.Val] = struct{}{}
		return find(node.Left, k) || find(node.Right, k)
	}
	return find(root, x)
}

func main() {
	node := &TreeNode{Val: 1}
	println(findTarget(node, 2))
}
