package sum_range

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestSumRange(t *testing.T) {
	testCases := []struct {
		i        int
		j        int
		expected int
	}{
		{
			i:        0,
			j:        2,
			expected: 1,
		},
		{
			i:        2,
			j:        5,
			expected: -1,
		},
		{
			i:        0,
			j:        5,
			expected: -3,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d %d", tc.i, tc.j)
		t.Run(name, func(t *testing.T) {
			nums := []int{-2, 0, 3, -5, 2, -1}
			obj := Constructor(nums)
			actual := obj.SumRange(tc.i, tc.j)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
