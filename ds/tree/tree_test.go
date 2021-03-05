package tree

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestTreeSize(t *testing.T) {
	var testCases = []struct {
		tree     *TreeNode
		expected int
	}{
		{
			tree:     &TreeNode{},
			expected: 1,
		},
		{
			tree:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			tree: &TreeNode{
				Val:   1,
				Left:  &TreeNode{},
				Right: &TreeNode{},
			},
			expected: 3,
		},
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{Val: 5},
			},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("Tree: %d", tc.tree.List())
		t.Run(name, func(t *testing.T) {
			actual := tc.tree.Size()
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
