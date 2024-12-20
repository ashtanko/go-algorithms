package remove_duplicate_letters

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

type testCase struct {
	s        string
	expected string
}

func TestRemoveDuplicateLetters(t *testing.T) {

	inputData := []testCase{
		{
			"bcabc",
			"abc",
		},
		{
			"cbacdcbc",
			"acdb",
		},
		{
			"abc",
			"abc",
		},
		{
			"cbc",
			"bc",
		},
	}

	testCases := []struct {
		strategy     func(string) string
		strategyName string
	}{
		{
			strategy:     removeDuplicateLetters,
			strategyName: "greedy",
		},
	}

	for _, tc := range testCases {
		for _, i := range inputData {
			name := fmt.Sprintf("%s %s", tc.strategyName, i.s)
			t.Run(name, func(t *testing.T) {
				actual := tc.strategy(i.s)
				utils.Checkf(t, is.DeepEqual(i.expected, actual), tc)
			})
		}
	}
}
