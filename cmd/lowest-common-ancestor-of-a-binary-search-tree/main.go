package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func main() {

}

//func findNodeInTree(node *TreeNode, val int) bool {
//	if node.Val == val {
//		return true
//	}
//	if node.Val > val {
//		return findNodeInTree(node.Left, val)
//	}
//	return findNodeInTree(node.Right, val)
//}
