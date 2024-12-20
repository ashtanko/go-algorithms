package increasing_triplet

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

type testCase struct {
	nums     []int
	expected bool
}

func TestIncreasingTriplet(t *testing.T) {

	inputData := []testCase{
		{
			[]int{1, 2, 3, 4, 5},
			true,
		},
		{
			[]int{5, 4, 3, 2, 1},
			false,
		},
		{
			[]int{2, 1, 5, 0, 4, 6},
			true,
		},
	}

	testCases := []struct {
		strategy     func(nums []int) bool
		strategyName string
	}{
		{
			strategy:     increasingTriplet,
			strategyName: "greedy",
		},
	}

	for _, tc := range testCases {
		for _, i := range inputData {
			name := fmt.Sprintf("%s %d", tc.strategyName, i.nums)
			t.Run(name, func(t *testing.T) {
				actual := tc.strategy(i.nums)
				utils.Checkf(t, is.DeepEqual(i.expected, actual), tc)
			})
		}
	}
}
