package xor_operation

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestXorOperation(t *testing.T) {
	testCases := []struct {
		n        int
		start    int
		expected int
	}{
		{
			n:        5,
			start:    0,
			expected: 8,
		},
		{
			n:        4,
			start:    3,
			expected: 8,
		},
		{
			n:        1,
			start:    7,
			expected: 7,
		},
		{
			n:        10,
			start:    5,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d %d", tc.n, tc.start)
		t.Run(name, func(t *testing.T) {
			actual := xorOperation(tc.n, tc.start)
			utils.Checkf(t, is.DeepEqual(actual, tc.expected), tc)
		})
	}
}
