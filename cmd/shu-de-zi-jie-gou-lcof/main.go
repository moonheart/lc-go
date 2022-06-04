package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return isSubStructure1(A, B, true)
}

func isSubStructure1(A *TreeNode, B *TreeNode, isBRoot bool) bool {
	if A == nil || B == nil {
		return false
	}
	if A.Val == B.Val {
		isLeftOk, isRightOk := false, false
		if B.Left != nil && A.Left != nil {
			isLeftOk = isSubStructure1(A.Left, B.Left, false)
		} else if B.Left == nil {
			isLeftOk = true
		}
		if B.Right != nil && A.Right != nil {
			isRightOk = isSubStructure1(A.Right, B.Right, false)
		} else if B.Right == nil {
			isRightOk = true
		}
		if isLeftOk && isRightOk {
			return true
		}
	}
	if !isBRoot {
		return false
	}
	return isSubStructure1(A.Left, B, true) || isSubStructure1(A.Right, B, true)
}

func main() {
	A := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   0,
			Left:  &TreeNode{Val: -4},
			Right: &TreeNode{Val: -3},
		},
		Right: &TreeNode{Val: 1},
	}
	B := &TreeNode{Val: 1, Left: &TreeNode{Val: -4}}
	fmt.Println(isSubStructure(A, B))
}
