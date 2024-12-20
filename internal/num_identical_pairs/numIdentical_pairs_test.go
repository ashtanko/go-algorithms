package num_identical_pairs

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumIdenticalPairs(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 2, 3, 1, 1, 3},
			expected: 4,
		},
		{
			nums:     []int{1, 1, 1, 1},
			expected: 6,
		},
		{
			nums:     []int{1, 2, 3},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.nums)
		t.Run(name, func(t *testing.T) {
			actual := numIdenticalPairs(tc.nums)
			utils.Checkf(t, is.DeepEqual(actual, tc.expected), tc)
		})
	}
}
