package word_break

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

type testCase struct {
	s        string
	wordDict []string
	expected bool
}

func TestWordBreak(t *testing.T) {

	inputData := []testCase{
		{
			"",
			[]string{"leet"},
			true,
		},
		{
			"",
			[]string{},
			true,
		},
		{
			"a",
			[]string{},
			false,
		},
		{
			"a",
			[]string{"a"},
			true,
		},
		{
			"leetcode",
			[]string{"leet", "code"},
			true,
		},
		{
			"applepenapple",
			[]string{"apple", "pen"},
			true,
		},
		{
			"catsandog",
			[]string{"cats", "dog", "sand", "and", "cat"},
			false,
		},
	}

	testCases := []struct {
		input        []testCase
		strategy     func(string, []string) bool
		strategyName string
	}{
		{
			input:        inputData,
			strategy:     wordBreakTopDown,
			strategyName: "top down",
		},
		{
			input:        inputData,
			strategy:     wordBreakMemo,
			strategyName: "top down memo",
		},
		{
			input:        inputData,
			strategy:     wordBreak1D,
			strategyName: "DP 1D",
		},
	}

	for _, tc := range testCases {
		for _, i := range tc.input {
			name := fmt.Sprintf("%s %s", i.s, tc.strategyName)
			t.Run(name, func(t *testing.T) {
				actual := tc.strategy(i.s, i.wordDict)
				utils.Checkf(t, is.DeepEqual(i.expected, actual), tc)
			})
		}
	}
}
