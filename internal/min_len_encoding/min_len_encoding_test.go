package min_len_encoding

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinLenEncoding(t *testing.T) {
	testCases := []struct {
		words    []string
		expected int
	}{
		{
			words:    []string{"time", "me", "bell"},
			expected: 10,
		},
		{
			words:    []string{"t"},
			expected: 2,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%s", tc.words)
		t.Run(name, func(t *testing.T) {
			actual := minimumLengthEncoding(tc.words)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
