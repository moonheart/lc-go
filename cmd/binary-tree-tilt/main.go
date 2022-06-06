package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	var dfs func(node *TreeNode) (sum int)
	tilt := 0
	dfs = func(node *TreeNode) (sum int) {
		if node == nil {
			return 0
		}
		L := dfs(node.Left)
		R := dfs(node.Right)
		tilt += abs(L - R)
		return L + R + node.Val
	}
	dfs(root)
	return tilt
}

func abs(int2 int) int {
	if int2 < 0 {
		return -int2
	}
	return int2
}

func main() {
	node := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   9,
			Right: &TreeNode{Val: 7},
		},
	}
	println(findTilt(node))
}
