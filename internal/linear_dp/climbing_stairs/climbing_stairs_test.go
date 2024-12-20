package climbing_stairs

import (
	"fmt"
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	testCases := []struct {
		num      int
		expected int
	}{
		{
			num:      0,
			expected: 1,
		},
		{
			num:      1,
			expected: 1,
		},
		{
			num:      2,
			expected: 2,
		},
		{
			num:      3,
			expected: 3,
		},
		{
			num:      5,
			expected: 8,
		},
		{
			num:      20,
			expected: 10946,
		},
		{
			num:      30,
			expected: 1346269,
		},
		{
			num:      40,
			expected: 165580141,
		},
		{
			num:      60,
			expected: 2504730781961,
		},
		{
			num:      70,
			expected: 308061521170129,
		},
		{
			num:      90,
			expected: 4660046610375530309,
		},
		{
			num:      100,
			expected: 1298777728820984005,
		},
		{
			num:      120,
			expected: 790376311979428689,
		},
		{
			num:      140,
			expected: 1233533820207347330,
		},
		{
			num:      160,
			expected: 9217463444206948445,
		},
		{
			num:      170,
			expected: 1812812405337582786,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.num)
		t.Run(name, func(t *testing.T) {
			actual := climbingStairsBottomUpDp(tc.num)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
