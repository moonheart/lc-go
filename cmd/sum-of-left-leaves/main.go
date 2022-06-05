package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	if root.Right != nil {
		ans += sumOfLeftLeaves(root.Right)
	}

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			ans += root.Left.Val
		} else {
			ans += sumOfLeftLeaves(root.Left)
		}
	}
	return
}

func main() {
	node := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	println(sumOfLeftLeaves(node))
}
