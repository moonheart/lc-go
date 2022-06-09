package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var preOrder func(node *TreeNode, arr *[]int)
	preOrder = func(node *TreeNode, arr *[]int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			*arr = append(*arr, node.Val)
			return
		}
		preOrder(node.Left, arr)
		preOrder(node.Right, arr)
	}

	var arr1 []int
	var arr2 []int
	preOrder(root1, &arr1)
	preOrder(root2, &arr2)

	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
