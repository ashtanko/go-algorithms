package running_sum

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestRunningSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{1, 3, 6, 10},
		},
		{
			nums:     []int{1, 1, 1, 1, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			nums:     []int{3, 1, 2, 10, 1},
			expected: []int{3, 4, 6, 16, 17},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.nums)
		t.Run(name, func(t *testing.T) {
			actual := runningSum(tc.nums)
			utils.Checkf(t, is.DeepEqual(actual, tc.expected), tc)
		})
	}
}
