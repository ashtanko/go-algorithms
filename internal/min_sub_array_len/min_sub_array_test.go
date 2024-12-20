package min_sub_array_len

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinSubArrayLength(t *testing.T) {
	testCases := []struct {
		target   int
		nums     []int
		expected int
	}{
		{
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2,
		},
		{
			target:   4,
			nums:     []int{1, 4, 4},
			expected: 1,
		},
		{
			target:   11,
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("target: %d nums: %d", tc.target, tc.nums)
		t.Run(name, func(t *testing.T) {
			actual := minSubArrayLen(tc.target, tc.nums)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
