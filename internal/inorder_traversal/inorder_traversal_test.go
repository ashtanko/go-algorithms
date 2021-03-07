package inorder_traversal

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/ds/tree"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	testCases := []struct {
		root     *tree.TreeNode
		expected []int
	}{
		{
			root: &tree.TreeNode{
				Val: 1,
				Right: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 3, 2},
		},
		{
			root: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
				},
			},
			expected: []int{2, 1},
		},
		{
			root: &tree.TreeNode{
				Val: 1,
				Right: &tree.TreeNode{
					Val: 2,
				},
			},
			expected: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.root.List())
		t.Run(name, func(t *testing.T) {
			actual := inorderTraversal(tc.root)
			utils.Checkf(t, is.DeepEqual(tc.expected, actual), tc)
		})
	}
}
