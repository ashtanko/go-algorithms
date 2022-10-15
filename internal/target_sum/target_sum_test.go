package target_sum

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestTargetSum(t *testing.T) {

	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{1, 1, 1, 1, 1},
			target:   3,
			expected: 5,
		},
		{
			nums:     []int{1},
			target:   1,
			expected: 1,
		},
	}

	implementations := []utils.Pair[string, TargetSum]{
		{
			Left: "backtracking",
			Right: &TargetSumValueBacktracking{
				TargetSumValueBase: &TargetSumValueBase{},
			},
		},
		{
			Left: "recursive",
			Right: &TargetSumValueRecursive{
				TargetSumValueBase: &TargetSumValueBase{},
			},
		},
		{
			Left: "memoization",
			Right: &TargetSumValueMemoization{
				TargetSumValueBase: &TargetSumValueBase{},
			},
		},
		{
			Left: "targetSumValueDPWithMaps",
			Right: &TargetSumValueDPWithMaps{
				TargetSumValueBase: &TargetSumValueBase{},
			},
		},
		{
			Left: "targetSum2DDP",
			Right: &TargetSumValue2DDP{
				TargetSumValueBase: &TargetSumValueBase{},
			},
		},
		{
			Left: "targetSumValueDPLinearSpace",
			Right: &TargetSumValueDPLinearSpace{
				TargetSumValueBase: &TargetSumValueBase{},
			},
		},
	}

	for _, i := range implementations {
		for _, tc := range testCases {
			name := fmt.Sprintf("%s arr: %d target: %d exp: %d", i.Left, tc.nums, tc.target, tc.expected)
			t.Run(name, func(t *testing.T) {
				actual := i.Right.findTargetSumWays(tc.nums, tc.target)
				utils.Checkf(t, is.DeepEqual(actual, tc.expected), tc)
			})
		}
	}
}
