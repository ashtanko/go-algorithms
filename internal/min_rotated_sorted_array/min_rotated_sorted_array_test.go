package min_rotated_sorted_array

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestFindMinRotatedSorted(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			nums:     []int{11, 13, 15, 17},
			expected: 11,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.nums)
		t.Run(name, func(t *testing.T) {
			actual := findMinRotatedSorted(tc.nums)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
