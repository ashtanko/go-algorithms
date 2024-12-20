package koko_eating_bananas

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinEatingSpeed(t *testing.T) {
	testCases := []struct {
		piles    []int
		h        int
		expected int
	}{
		{
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		{
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},
		{
			piles:    []int{30, 11, 23, 4, 20},
			h:        6,
			expected: 23,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("piles: %d h: %d exp: %d", tc.piles, tc.h, tc.expected)
		t.Run(name, func(t *testing.T) {
			actual := minEatingSpeed(tc.piles, tc.h)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
