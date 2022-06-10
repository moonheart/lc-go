package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numColor(root *TreeNode) int {
	var m = map[int]struct{}{}
	var n func(root *TreeNode)
	n = func(root *TreeNode) {
		if root == nil {
			return
		}
		m[root.Val] = struct{}{}
		n(root.Left)
		n(root.Right)
	}
	n(root)
	return len(m)
}

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 1},
		},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}},
	}
	println(numColor(node))
}
