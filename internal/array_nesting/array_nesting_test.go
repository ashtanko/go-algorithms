package array_nesting

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestArrayNesting(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{5, 4, 0, 3, 1, 6, 2}, 4},
		{[]int{0, 1, 2, 3, 4, 5, 6}, 1},
		{[]int{0, 1, 2, 3, 4, 5, 5}, 1},
		{[]int{0, 0, 0, 0, 0, 0, 0}, 1},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.nums)
		t.Run(name, func(t *testing.T) {
			actual := arrayNesting(tc.nums)
			utils.Checkf(t, is.DeepEqual(tc.expected, actual), tc)
		})
	}
}
