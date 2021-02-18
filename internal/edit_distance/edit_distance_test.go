package edit_distance

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestEditDistance(t *testing.T) {
	testCases := []struct {
		word1    string
		word2    string
		expected int
	}{
		{
			word1:    "horse",
			word2:    "ros",
			expected: 3,
		},
		{
			word1:    "intention",
			word2:    "execution",
			expected: 5,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%s %s", tc.word1, tc.word2)
		t.Run(name, func(t *testing.T) {
			actual := minDistance(tc.word1, tc.word2)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
