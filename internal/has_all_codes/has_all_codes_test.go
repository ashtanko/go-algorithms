package has_all_codes

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestHasAllCodes(t *testing.T) {
	testCases := []struct {
		s        string
		k        int
		expected bool
	}{
		{
			"00110110",
			2,
			true,
		},
		{
			"00110",
			2,
			true,
		},
		{
			"0110",
			1,
			true,
		},
		{
			"0110",
			2,
			false,
		},
		{
			"0000000001011100",
			4,
			false,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("s: %s k: %d", tc.s, tc.k)
		t.Run(name, func(t *testing.T) {
			actual := hasAllCodes(tc.s, tc.k)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
