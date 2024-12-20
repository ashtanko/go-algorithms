package sort_characters_by_frequency

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencySort(t *testing.T) {

	testCases := []struct {
		s        string
		expected string
	}{
		{
			s:        "tree",
			expected: "eetr",
		},
		{
			s:        "cccaaa",
			expected: "cccaaa",
		},
		{
			s:        "Aabb",
			expected: "bbaA",
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("Test case: %s", tc.s)
		t.Run(name, func(t *testing.T) {
			actual := frequencySort(tc.s)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
