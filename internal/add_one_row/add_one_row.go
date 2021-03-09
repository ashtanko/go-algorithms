package add_one_row

import (
	"github.com/ashtanko/go-algorithms/ds/tree"
)

func addOneRow(root *tree.TreeNode, v, d int) *tree.TreeNode {
	if d == 1 {
		return &tree.TreeNode{
			Val:   v,
			Left:  root,
			Right: nil,
		}
	}
	return insertToTree(root, v, d-1, 1)
}

func insertToTree(root *tree.TreeNode, value, depth, n int) *tree.TreeNode {
	if root == nil {
		return nil
	}

	if depth == n {
		root.Left = &tree.TreeNode{
			Val:   value,
			Left:  root.Left,
			Right: nil,
		}

		root.Right = &tree.TreeNode{
			Val:   value,
			Left:  nil,
			Right: root.Right,
		}
	}
	root.Left = insertToTree(root.Left, value, depth, n+1)
	root.Right = insertToTree(root.Right, value, depth, n+1)
	return root
}
