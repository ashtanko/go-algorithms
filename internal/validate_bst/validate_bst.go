package validate_bst

import (
	"github.com/ashtanko/go-algorithms/internal/ds/tree"
)

func isValidBST(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	res, v := true, root.Val
	doValidateBST(&res, root.Left, &v, nil)
	doValidateBST(&res, root.Right, nil, &v)
	return res
}

func doValidateBST(res *bool, node *tree.TreeNode, upperBound, lowerBound *int) {
	if node == nil || !*res {
		return
	}
	if (upperBound != nil && node.Val >= *upperBound) ||
		(lowerBound != nil && node.Val <= *lowerBound) {
		*res = false
		return
	}

	v := node.Val
	doValidateBST(res, node.Left, &v, lowerBound)
	doValidateBST(res, node.Right, upperBound, &v)
}
