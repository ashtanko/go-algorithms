package int_to_roman

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		num      int
		expected string
	}{
		{
			3,
			"III",
		},
		{
			4,
			"IV",
		},
		{
			9,
			"IX",
		},
		{
			58,
			"LVIII",
		},
		{
			1994,
			"MCMXCIV",
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.num)
		t.Run(name, func(t *testing.T) {
			actual := intToRoman(tc.num)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
