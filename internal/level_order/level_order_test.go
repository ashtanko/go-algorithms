package level_order

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/internal/ds/tree"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestLevelOrder(t *testing.T) {
	testCases := []struct {
		tree     *tree.TreeNode
		expected [][]int
	}{
		{
			tree:     nil,
			expected: nil,
		},
		{
			tree:     &tree.TreeNode{},
			expected: [][]int{{0}},
		},
		{
			tree: &tree.TreeNode{
				Val: 3,
				Left: &tree.TreeNode{
					Val: 9,
				},
				Right: &tree.TreeNode{
					Val: 20,
					Left: &tree.TreeNode{
						Val: 15,
					},
					Right: &tree.TreeNode{
						Val: 7,
					},
				},
			},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.tree.List())
		t.Run(name, func(t *testing.T) {
			actual := levelOrder(tc.tree)
			utils.Checkf(t, is.DeepEqual(tc.expected, actual), tc)
		})
	}
}
