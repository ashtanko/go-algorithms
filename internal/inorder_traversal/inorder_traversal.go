package inorder_traversal

import "github.com/ashtanko/go-algorithms/ds/tree"

func inorderTraversal(root *tree.TreeNode) []int {
	var (
		stack  []*tree.TreeNode
		result []int
		cur    = root
	)

	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right
	}

	return result
}
