package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := len(nums) / 2
	node := &TreeNode{Val: nums[m]}
	if m == 0 {
		return node
	}
	node.Left = sortedArrayToBST(nums[:m])
	node.Right = sortedArrayToBST(nums[m+1:])
	return node
}

func main() {
	bst := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	marshal, _ := json.MarshalIndent(bst, "", "  ")
	fmt.Print(string(marshal))
}
