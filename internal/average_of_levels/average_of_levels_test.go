package average_of_levels

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/internal/ds/tree"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	testCases := []struct {
		tree     *tree.TreeNode
		expected []float64
	}{
		{
			tree:     nil,
			expected: nil,
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
			expected: []float64{3, 14.5, 11},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.tree.List())
		t.Run(name, func(t *testing.T) {
			actual := averageOfLevels(tc.tree)
			utils.Checkf(t, is.DeepEqual(tc.expected, actual), tc)
		})
	}
}
