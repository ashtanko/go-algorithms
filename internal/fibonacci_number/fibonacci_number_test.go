package fibonacci_number

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

type fib struct {
	n        int
	expected int
}

func TestFibonacciNumber(t *testing.T) {
	fibs := []fib{
		{
			n:        0,
			expected: 0,
		},
		{
			n:        1,
			expected: 1,
		},
		{
			n:        2,
			expected: 1,
		},
		{
			n:        3,
			expected: 2,
		},
		{
			n:        4,
			expected: 3,
		},
	}

	testCases := []struct {
		fibs    []fib
		perform func(int) int
	}{
		{
			fibs:    fibs,
			perform: recursive,
		},
		{
			fibs:    fibs,
			perform: bottomUp,
		},
		{
			fibs:    fibs,
			perform: topDown,
		}, {
			fibs:    fibs,
			perform: fibMath,
		},
	}

	for _, tc := range testCases {
		for _, f := range tc.fibs {
			name := fmt.Sprintf("%d", f.n)
			t.Run(name, func(t *testing.T) {
				actual := tc.perform(f.n)
				utils.Checkf(t, is.Equal(f.expected, actual), tc)
			})
		}
	}
}
