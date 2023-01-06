package permutation_sequence

import (
	"fmt"
	"testing"
)

var (
	tests = []struct {
		n        int
		k        int
		expected string
	}{
		{3, 3, "213"},
		{4, 9, "2314"},
		{3, 1, "123"},
	}

	task = []func(n int, k int) string{getPermutation}
)

func TestGetPermutation(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			test := test
			t.Run(fmt.Sprint(test.n, test.expected), func(t *testing.T) {
				t.Parallel()
				if have := fn(test.n, test.k); have != test.expected {
					t.Errorf(`input: %+vexpected: %+v`, test.expected, have)
				}
			})
		}
	}
}
