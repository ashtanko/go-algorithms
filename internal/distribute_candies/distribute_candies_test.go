package distribute_candies

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestDistributeCandies(t *testing.T) {
	testCases := []struct {
		candyType []int
		expected  int
	}{
		{
			[]int{1, 1, 2, 2, 3, 3},
			3,
		},
		{
			[]int{1, 1, 2, 3},
			2,
		},
		{
			[]int{6, 6, 6, 6},
			1,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.candyType)
		t.Run(name, func(t *testing.T) {
			actual := distributeCandies(tc.candyType)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
