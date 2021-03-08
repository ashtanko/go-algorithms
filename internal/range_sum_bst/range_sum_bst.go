package range_sum_bst

import "github.com/ashtanko/go-algorithms/ds/tree"

func rangeSumBST(root *tree.TreeNode, low int, high int) int {
	return tree.Dfs(root, low, high)
}
