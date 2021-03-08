package range_sum_bst

import "github.com/ashtanko/go-algorithms/ds/tree"

func rangeSumBST(root *tree.TreeNode, low int, high int) int {
	ans := 0
	tree.Dfs(root, low, high, &ans)
	return ans
}
