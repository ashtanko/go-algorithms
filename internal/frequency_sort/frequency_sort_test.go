package frequency_sort

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 1, 2, 2, 2, 3},
			expected: []int{3, 1, 1, 2, 2, 2},
		},
		{
			nums:     []int{2, 3, 1, 3, 2},
			expected: []int{1, 3, 3, 2, 2},
		},
		{
			nums:     []int{-1, 1, -6, 4, 5, -6, 1, 4, 1},
			expected: []int{5, -1, 4, 4, -6, -6, 1, 1, 1},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.nums)
		t.Run(name, func(t *testing.T) {
			actual := frequencySort(tc.nums)
			utils.Checkf(t, is.DeepEqual(tc.expected, actual), tc)
		})
	}
}
