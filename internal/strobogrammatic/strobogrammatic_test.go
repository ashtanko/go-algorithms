package strobogrammatic

import (
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestStrobogrammaticNumber(t *testing.T) {
	testCases := []struct {
		num      string
		expected bool
	}{
		{
			"69",
			true,
		},
		{
			"88",
			true,
		},
		{
			"962",
			false,
		},
		{
			"1",
			true,
		},
	}

	for _, tc := range testCases {
		name := tc.num
		t.Run(name, func(t *testing.T) {
			actual := isStrobogrammatic(tc.num)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
