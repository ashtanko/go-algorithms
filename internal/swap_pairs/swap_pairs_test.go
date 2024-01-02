package swap_pairs

import (
	"fmt"
	. "github.com/ashtanko/go-algorithms/internal/ds/linked_list"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestSwapPairs(t *testing.T) {

	testCases := []struct {
		head     []int
		expected []int
	}{
		{
			head:     []int{1, 2, 3, 4},
			expected: []int{2, 1, 4, 3},
		},
		{
			head:     []int{},
			expected: []int{},
		},
		{
			head:     []int{1},
			expected: []int{1},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.head)
		t.Run(name, func(t *testing.T) {
			actual := swapPairs(FromSlice(tc.head)).ToSlice()
			utils.Checkf(t, is.DeepEqual(actual, tc.expected), tc)
		})
	}
}
