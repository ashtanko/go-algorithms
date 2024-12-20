package max_profit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitFee(t *testing.T) {
	testCases := []struct {
		prices   []int
		fee      int
		expected int
	}{
		{
			prices:   []int{1, 3, 2, 8, 4, 9},
			fee:      2,
			expected: 8,
		},
		{
			prices:   []int{1, 3, 7, 5, 10, 3},
			fee:      3,
			expected: 6,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.prices)
		t.Run(name, func(t *testing.T) {
			actual := maxProfitFee(tc.prices, tc.fee)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
