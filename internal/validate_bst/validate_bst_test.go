package validate_bst

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/internal/ds/tree"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestValidateBinarySearchTree(t *testing.T) {
	testCases := []struct {
		root     *tree.TreeNode
		expected bool
	}{
		{
			root: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 1,
				},
				Right: &tree.TreeNode{
					Val: 3,
				},
			},
			expected: true,
		},
		{
			root: &tree.TreeNode{
				Val: 5,
				Left: &tree.TreeNode{
					Val: 1,
				},
				Right: &tree.TreeNode{
					Val: 4,
					Left: &tree.TreeNode{
						Val: 3,
					},
					Right: &tree.TreeNode{
						Val: 6,
					},
				},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.root.List())
		t.Run(name, func(t *testing.T) {
			actual := isValidBST(tc.root)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
