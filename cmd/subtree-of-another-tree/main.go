package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	lb = 999
	mb = 999999
	rb = 999999999
)

func hash(node *TreeNode) int {
	if node == nil {
		return 1
	}
	node.Val = lb*hash(node.Left) + rb*hash(node.Right) + mb*node.Val
	return node.Val
}

func dfs(node *TreeNode, val int) bool {
	if node == nil {
		return false
	}
	return node.Val == val || dfs(node.Left, val) || dfs(node.Right, val)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	_ = hash(root)
	subH := hash(subRoot)
	return dfs(root, subH)
}

func isSubtree1(root *TreeNode, subRoot *TreeNode) bool {
	var isSub func(node *TreeNode, subNode *TreeNode, isMatching bool) bool
	isSub = func(node *TreeNode, subNode *TreeNode, isMatching bool) bool {
		if node == nil && subNode == nil {
			return true
		}
		if node == nil || subNode == nil {
			return false
		}
		if node.Val == subNode.Val && isSub(node.Left, subNode.Left, true) && isSub(node.Right, subNode.Right, true) {
			return true
		} else if isMatching {
			return false
		}
		return isSub(node.Left, subNode, false) || isSub(node.Right, subNode, false)
	}
	return isSub(root, subRoot, false)
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 7},
			},
		},
	}
	subroot := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
	println(isSubtree(root, subroot))
}
