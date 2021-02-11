package maximum_69_number

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximum69Number(t *testing.T) {
	testCases := []struct {
		num      int
		expected int
	}{
		{
			num:      6,
			expected: 9,
		},
		{
			num:      9,
			expected: 9,
		},
		{
			num:      9669,
			expected: 9969,
		},
		{
			num:      9996,
			expected: 9999,
		},
		{
			num:      9999,
			expected: 9999,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.num)
		t.Run(name, func(t *testing.T) {
			actual := maximum69Number(tc.num)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
