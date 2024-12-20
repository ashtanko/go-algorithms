package lonely_nodes

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/internal/ds/tree"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestLonelyNodes(t *testing.T) {
	testCases := []struct {
		tree     *tree.TreeNode
		expected []int
	}{
		{
			tree: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
					Right: &tree.TreeNode{
						Val: 4,
					},
				},
				Right: &tree.TreeNode{
					Val: 3,
				},
			},
			expected: []int{4},
		},
		{
			tree: &tree.TreeNode{
				Val: 7,
				Left: &tree.TreeNode{
					Val: 1,
					Left: &tree.TreeNode{
						Val: 6,
					},
				},
				Right: &tree.TreeNode{
					Val: 4,
					Left: &tree.TreeNode{
						Val: 5,
					},
					Right: &tree.TreeNode{
						Val: 3,
						Right: &tree.TreeNode{
							Val: 2,
						},
					},
				},
			},
			expected: []int{6, 2},
		},
		{
			tree: &tree.TreeNode{
				Val: 11,
				Left: &tree.TreeNode{
					Val: 99,
					Left: &tree.TreeNode{
						Val: 77,
						Left: &tree.TreeNode{
							Val: 55,
							Left: &tree.TreeNode{
								Val: 33,
							},
						},
					},
				},
				Right: &tree.TreeNode{
					Val: 88,
					Right: &tree.TreeNode{
						Val: 66,
						Right: &tree.TreeNode{
							Val: 44,
							Right: &tree.TreeNode{
								Val: 22,
							},
						},
					},
				},
			},
			expected: []int{77, 55, 33, 66, 44, 22},
		},
		{
			tree: &tree.TreeNode{
				Val: 197,
			},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.tree.List())
		t.Run(name, func(t *testing.T) {
			actual := getLonelyNodes(tc.tree)
			utils.Checkf(t, is.DeepEqual(tc.expected, actual), tc)
		})
	}
}
