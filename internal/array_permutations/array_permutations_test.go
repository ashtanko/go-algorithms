package array_permutations

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestArrayPermutations(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			[]int{0, 1},
			[][]int{{0, 1}, {1, 0}},
		},
		{
			[]int{1},
			[][]int{{1}},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.nums)
		t.Run(name, func(t *testing.T) {
			actual := permute(tc.nums)
			utils.Checkf(t, is.DeepEqual(tc.expected, actual), tc)
		})
	}
}
