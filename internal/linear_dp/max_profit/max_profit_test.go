package max_profit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		prices   []int
		expected int
	}{
		{
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.prices)
		t.Run(name, func(t *testing.T) {
			actual := maxProfit(tc.prices)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
