package lonely_nodes

import (
	"github.com/ashtanko/go-algorithms/internal/ds/tree"
)

func getLonelyNodes(root *tree.TreeNode) []int {
	return dfs(root, nil)
}

func dfs(root *tree.TreeNode, parent *tree.TreeNode) []int {
	lonely := make([]int, 0)
	if root == nil {
		return lonely
	}

	if parent != nil && (parent.Left == nil || parent.Right == nil) {
		lonely = append(lonely, root.Val)
	}

	lonely = append(lonely, dfs(root.Left, root)...)
	lonely = append(lonely, dfs(root.Right, root)...)

	return lonely
}
