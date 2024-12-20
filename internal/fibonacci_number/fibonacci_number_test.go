package fibonacci_number

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

type fib struct {
	n        int
	expected int
}

type strategy struct {
	name     string
	solution func(int) int
}

var (
	tests = []fib{
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

	tasks = []strategy{
		{
			"recursive",
			recursive,
		},
		{
			"bottom up",
			bottomUp,
		},
		{
			"top down",
			topDown,
		},
		{
			"math",
			fibMath,
		},
	}
)

func TestFibonacciNumber(t *testing.T) {
	for _, fn := range tasks {
		for _, f := range tests {
			name := fmt.Sprintf("%s %d exp: %d", fn.name, f.n, f.expected)
			t.Run(name, func(t *testing.T) {
				actual := fn.solution(f.n)
				utils.Checkf(t, is.Equal(f.expected, actual), t)
			})
		}
	}
}
