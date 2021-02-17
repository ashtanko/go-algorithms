package min_cost_climbing_stairs

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	testCases := []struct {
		cost     []int
		expected int
	}{
		{
			cost:     []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			cost:     []int{10, 15, 20},
			expected: 15,
		},
		{
			cost:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.cost)
		t.Run(name, func(t *testing.T) {
			actual := minCostClimbingStairs(tc.cost)
			utils.Checkf(t, is.DeepEqual(actual, tc.expected), tc)
		})
	}
}
