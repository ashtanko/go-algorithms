package wiggle_max_length

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestWiggleMaxLength(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 7, 4, 9, 2, 5},
			6,
		},
		{
			[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			2,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.nums)
		t.Run(name, func(t *testing.T) {
			actual := wiggleMaxLength(tc.nums)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
