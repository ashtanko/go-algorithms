package sum_numbers

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/internal/ds/tree"
	"testing"
)

var (
	tests = []struct {
		root     *tree.TreeNode
		expected int
	}{
		{
			root: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
				},
				Right: &tree.TreeNode{
					Val: 3,
				},
			},
			expected: 25,
		},
		{
			root: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 9,
					Right: &tree.TreeNode{
						Val: 1,
					},
					Left: &tree.TreeNode{
						Val: 5,
					},
				},
				Right: &tree.TreeNode{
					Val: 0,
				},
			},
			expected: 1026,
		},
	}

	task = []func(root *tree.TreeNode) int{sumNumbers, sumNumbersDFS}
)

func TestSumNumbers(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			t.Run(fmt.Sprint(test.root, test.expected), func(t *testing.T) {
				if have := fn(test.root); have != test.expected {
					t.Errorf(`input: %+vexpected: %+v`, test.expected, have)
				}
			})
		}
	}
}
