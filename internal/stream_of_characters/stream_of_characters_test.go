package stream_of_characters

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestStreamChecker(t *testing.T) {

	testCases := []struct {
		letter   byte
		expected bool
	}{
		{
			letter:   'a',
			expected: false,
		},
		{
			letter:   'b',
			expected: false,
		},
		{
			letter:   'c',
			expected: false,
		},
		{
			letter:   'd',
			expected: true,
		},
		{
			letter:   'e',
			expected: false,
		},
		{
			letter:   'f',
			expected: true,
		},
		{
			letter:   'g',
			expected: false,
		},
		{
			letter:   'h',
			expected: false,
		},
		{
			letter:   'i',
			expected: false,
		},
		{
			letter:   'k',
			expected: false,
		},
		{
			letter:   'l',
			expected: true,
		},
	}

	words := []string{"cd", "f", "kl"}
	streamChecker := Constructor(words)

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.letter)
		t.Run(name, func(t *testing.T) {
			utils.Checkf(t, is.DeepEqual(streamChecker.Query(tc.letter), tc.expected), tc)
		})
	}
}
