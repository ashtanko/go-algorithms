package first_occurrence_in_string

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestFirstOccurrenceInString(t *testing.T) {
	testCases := []struct {
		haystack string
		needle   string
		expected int
	}{
		{
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
			expected: -1,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%s %s", tc.haystack, tc.needle)
		t.Run(name, func(t *testing.T) {
			actual := strStr(tc.haystack, tc.needle)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
