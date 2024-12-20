package same_tree

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/internal/ds/tree"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

type testCase struct {
	p        *tree.TreeNode
	q        *tree.TreeNode
	expected bool
}

func TestSameTree(t *testing.T) {

	inputData := []testCase{
		{
			p: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
				},
				Right: &tree.TreeNode{
					Val: 3,
				},
			},
			q: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
				},
				Right: &tree.TreeNode{
					Val: 3,
				},
			},
			expected: true,
		},
		{
			p: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
				},
				Right: &tree.TreeNode{
					Val: 3,
				},
			},
			q: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
				},
				Right: &tree.TreeNode{
					Val: 4,
				},
			},
			expected: false,
		},
		{
			p: &tree.TreeNode{
				Val: 1,
			},
			q: &tree.TreeNode{
				Val: 1,
			},
			expected: true,
		},
		{
			p: &tree.TreeNode{
				Val: 1,
			},
			q:        &tree.TreeNode{},
			expected: false,
		},
		{
			p: &tree.TreeNode{},
			q: &tree.TreeNode{
				Val: 1,
			},
			expected: false,
		},
		{
			p:        &tree.TreeNode{},
			q:        &tree.TreeNode{},
			expected: true,
		},
	}

	testCases := []struct {
		strategy     func(p *tree.TreeNode, q *tree.TreeNode) bool
		strategyName string
	}{
		{
			strategy:     isSameTreeRecursive,
			strategyName: "recursive",
		},
		{
			strategy:     isSameTreeBFS,
			strategyName: "BFS",
		},
	}

	for _, tc := range testCases {
		for _, i := range inputData {
			name := fmt.Sprintf("%s %d %d", tc.strategyName, i.p.List(), i.q.List())
			t.Run(name, func(t *testing.T) {
				actual := tc.strategy(i.p, i.q)
				utils.Checkf(t, is.DeepEqual(i.expected, actual), tc)
			})
		}
	}
}
