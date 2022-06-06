package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func diameterOfBinaryTree(root *TreeNode) int {
	var depth func(node *TreeNode) int
	max := 0
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		dL := depth(node.Left)
		dR := depth(node.Right)
		if dL+dR > max {
			max = dL + dR
		}
		if dL > dR {
			return dL + 1
		}
		return dR + 1
	}
	depth(root)

	return max
}

func main() {
	node := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3},
	}
	println(diameterOfBinaryTree(node))
}
