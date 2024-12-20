package is_subsequence

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsSubsequence(t *testing.T) {
	testCases := []struct {
		source   string
		target   string
		expected bool
	}{
		{
			source:   "abc",
			target:   "ahbgdc",
			expected: true,
		},
		{
			source:   "axc",
			target:   "ahbgdc",
			expected: false,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("source: %s target: %s", tc.source, tc.target)
		t.Run(name, func(t *testing.T) {
			actual := isSubsequence(tc.source, tc.target)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
