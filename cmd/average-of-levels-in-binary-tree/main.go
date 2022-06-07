package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		tq := q
		q = nil
		sum := 0
		for _, n := range tq {
			sum += n.Val
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		res = append(res, float64(sum)/float64(len(tq)))
	}
	return res
}
