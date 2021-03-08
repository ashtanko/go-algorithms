package tree

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	"gotest.tools/v3/assert"
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

func TestDfs(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	ans := 0
	Dfs(tree, 7, 15, &ans)
	assert.Equal(t, 32, ans)
}
