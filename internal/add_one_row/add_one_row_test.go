package add_one_row

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/ds/tree"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestAddOneRowToTree(t *testing.T) {
	testCases := []struct {
		root     *tree.TreeNode
		v, d     int
		expected *tree.TreeNode
	}{
		{
			root: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 3,
					},
					Right: &tree.TreeNode{
						Val: 1,
					},
				},
				Right: &tree.TreeNode{
					Val: 6,
					Left: &tree.TreeNode{
						Val: 5,
					},
				},
			},
			v: 1,
			d: 2,
			expected: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 1,
					Left: &tree.TreeNode{
						Val: 2,
						Left: &tree.TreeNode{
							Val: 3,
						},
						Right: &tree.TreeNode{
							Val: 1,
						},
					},
				},
				Right: &tree.TreeNode{
					Val: 1,
					Right: &tree.TreeNode{
						Val: 6,
						Left: &tree.TreeNode{
							Val: 5,
						},
					},
				},
			},
		},
		{
			root: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 3,
					},
					Right: &tree.TreeNode{
						Val: 1,
					},
				},
			},
			v: 1,
			d: 3,
			expected: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 1,
						Left: &tree.TreeNode{
							Val: 3,
						},
					},
					Right: &tree.TreeNode{
						Val: 1,
						Right: &tree.TreeNode{
							Val: 1,
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("root: %d v: %d d: %d", tc.root.List(), tc.v, tc.d)
		t.Run(name, func(t *testing.T) {
			actual := addOneRow(tc.root, tc.v, tc.d)
			fmt.Println("  actual: ", actual.List())
			fmt.Println("expected: ", tc.expected.List())
			utils.Checkf(t, is.DeepEqual(tc.expected.List(), actual.List()), tc)
		})
	}
}
