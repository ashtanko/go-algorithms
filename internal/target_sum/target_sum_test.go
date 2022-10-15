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

	var backtracking TargetSum
	backtracking = &TargetSumValueBacktracking{
		TargetSumValueBase: &TargetSumValueBase{},
	}

	var recursive TargetSum
	recursive = &TargetSumValueRecursive{
		TargetSumValueBase: &TargetSumValueBase{},
	}

	var memoization TargetSum
	memoization = &TargetSumValueMemoization{
		TargetSumValueBase: &TargetSumValueBase{},
	}

	var targetSum2DDP TargetSum
	targetSum2DDP = &TargetSumValue2DDP{
		TargetSumValueBase: &TargetSumValueBase{},
	}

	var targetSumValueDPLinearSpace TargetSum
	targetSumValueDPLinearSpace = &TargetSumValueDPLinearSpace{
		TargetSumValueBase: &TargetSumValueBase{},
	}

	var targetSumValueDPWithMaps TargetSum
	targetSumValueDPWithMaps = &TargetSumValueDPWithMaps{
		TargetSumValueBase: &TargetSumValueBase{},
	}

	implementations := []utils.Pair[string, TargetSum]{
		{
			Left:  "backtracking",
			Right: backtracking,
		},
		{
			Left:  "recursive",
			Right: recursive,
		},
		{
			Left:  "memoization",
			Right: memoization,
		},
		{
			Left:  "targetSumValueDPWithMaps",
			Right: targetSumValueDPWithMaps,
		},
		{
			Left:  "targetSum2DDP",
			Right: targetSum2DDP,
		},
		{
			Left:  "targetSumValueDPLinearSpace",
			Right: targetSumValueDPLinearSpace,
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
