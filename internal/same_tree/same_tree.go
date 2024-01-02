package same_tree

import (
	"github.com/ashtanko/go-algorithms/internal/ds/tree"
)

// 100. Same Tree
// https://leetcode.com/problems/same-tree/description/
func isSameTreeRecursive(p *tree.TreeNode, q *tree.TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return isSameTreeRecursive(p.Left, q.Left) && isSameTreeRecursive(p.Right, q.Right)
}

func isSameTreeBFS(p *tree.TreeNode, q *tree.TreeNode) bool {
	qP := []*tree.TreeNode{p}
	qQ := []*tree.TreeNode{q}

	for len(qP) != 0 && len(qQ) != 0 {
		pNode := qP[0]
		qP = qP[1:]

		qNode := qQ[0]
		qQ = qQ[1:]

		if pNode == nil && qNode == nil {
			continue
		}
		if pNode == nil && qNode != nil || pNode != nil && qNode == nil {
			return false
		}
		if pNode.Val != qNode.Val {
			return false
		}

		qP = append(qP, pNode.Left, pNode.Right)
		qQ = append(qQ, qNode.Left, qNode.Right)
	}

	if len(qP) == 0 && len(qQ) == 0 {
		return true
	}
	return false
}
