package range_sum_bst

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/ds/tree"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestRangeSumBST(t *testing.T) {
	testCases := []struct {
		root     *tree.TreeNode
		low      int
		high     int
		expected int
	}{
		{
			root: &tree.TreeNode{
				Val: 10,
				Left: &tree.TreeNode{
					Val: 5,
					Left: &tree.TreeNode{
						Val: 3,
					},
					Right: &tree.TreeNode{
						Val: 7,
					},
				},
				Right: &tree.TreeNode{
					Val: 15,
					Right: &tree.TreeNode{
						Val: 18,
					},
				},
			},
			low:      7,
			high:     15,
			expected: 32,
		},

		{
			root: &tree.TreeNode{
				Val: 10,
				Left: &tree.TreeNode{
					Val: 5,
					Left: &tree.TreeNode{
						Val: 3,
						Left: &tree.TreeNode{
							Val: 1,
						},
					},
					Right: &tree.TreeNode{
						Val: 7,
						Left: &tree.TreeNode{
							Val: 6,
						},
					},
				},
				Right: &tree.TreeNode{
					Val: 15,
					Left: &tree.TreeNode{
						Val: 13,
					},
					Right: &tree.TreeNode{
						Val: 18,
					},
				},
			},
			low:      6,
			high:     10,
			expected: 23,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("tree: %d low: %d hight: %d", tc.root.List(), tc.low, tc.high)
		t.Run(name, func(t *testing.T) {
			actual := rangeSumBST(tc.root, tc.low, tc.high)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
