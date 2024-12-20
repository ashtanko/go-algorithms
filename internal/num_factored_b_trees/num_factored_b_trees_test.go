package num_factored_b_trees

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestBinaryTreesWithFactors(t *testing.T) {
	testCases := []struct {
		arr      []int
		expected int
	}{
		{
			[]int{2, 4},
			3,
		},
		{
			[]int{2, 4, 5, 10},
			7,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.arr)
		t.Run(name, func(t *testing.T) {
			actual := numFactoredBinaryTrees(tc.arr)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
