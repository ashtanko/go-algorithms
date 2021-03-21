package reordered_power_of_two

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestReorderedPowerOfTwo(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{
			-1,
			false,
		},
		{
			0,
			false,
		},
		{
			1,
			true,
		},
		{
			10,
			false,
		},
		{
			16,
			true,
		},
		{
			24,
			false,
		},
		{
			46,
			true,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.n)
		t.Run(name, func(t *testing.T) {
			actual := reorderedPowerOf2(tc.n)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
