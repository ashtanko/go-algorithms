package check_array

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestCheckArray(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{},
			expected: true,
		},
		{
			nums:     []int{3, 4, 5, 1, 2},
			expected: true,
		},
		{
			nums:     []int{3, 1, 2},
			expected: true,
		},
		{
			nums:     []int{2, 1, 3, 4},
			expected: false,
		},
		{
			nums:     []int{1, 2, 3},
			expected: true,
		},
		{
			nums:     []int{1, 1, 1},
			expected: true,
		},
		{
			nums:     []int{2, 1},
			expected: true,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.nums)
		t.Run(name, func(t *testing.T) {
			actual := check(tc.nums)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
