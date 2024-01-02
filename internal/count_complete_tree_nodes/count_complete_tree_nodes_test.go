package count_complete_tree_nodes

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/internal/ds/tree"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestCountNodes(t *testing.T) {
	testCases := []struct {
		root     *tree.TreeNode
		expected int
	}{
		{
			root: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 4,
					},
					Right: &tree.TreeNode{
						Val: 5,
					},
				},
				Right: &tree.TreeNode{
					Val: 3,
					Left: &tree.TreeNode{
						Val: 5,
					},
				},
			},
			expected: 6,
		},
		{
			root: &tree.TreeNode{
				Val: 4,
			},
			expected: 1,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("root: %d", tc.root.List())
		t.Run(name, func(t *testing.T) {
			actual := countNodes(tc.root)
			fmt.Println("  actual: ", actual)
			fmt.Println("expected: ", tc.expected)
			utils.Checkf(t, is.DeepEqual(tc.expected, actual), tc)
		})
	}
}
