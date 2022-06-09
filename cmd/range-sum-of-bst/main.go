package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) (res int) {
	if root == nil {
		return
	}
	if root.Val > high {
		res += rangeSumBST(root.Left, low, high)
	} else if root.Val < low {
		res += rangeSumBST(root.Right, low, high)
	} else {
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}
	return
}
