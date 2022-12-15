package longest_common_subsequence

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	testCases := []struct {
		text1    string
		text2    string
		expected int
	}{
		{
			"",
			"",
			0,
		},
		{
			"abcde",
			"ace",
			3,
		},
		{
			"abc",
			"abc",
			3,
		},
		{
			"abcef",
			"abcf",
			4,
		},
		{
			"abc",
			"def",
			0,
		},
		{
			"",
			"",
			0,
		},
		{
			"q",
			"",
			0,
		},
		{
			"",
			"a",
			0,
		},
		{
			"ab",
			"abc",
			2,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%s %s", tc.text1, tc.text2)
		t.Run(name, func(t *testing.T) {
			actual := longestCommonSubsequence(tc.text1, tc.text2)
			utils.Checkf(t, is.DeepEqual(tc.expected, actual), tc)
		})
	}
}
