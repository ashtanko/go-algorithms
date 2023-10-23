package power_of_four

import (
	"fmt"
	"testing"
)

var (
	tests = []struct {
		n        int
		expected bool
	}{
		{16, true},
		{5, false},
		{1, true},
	}

	task = []func(n int) bool{isPowerOfFourBruteForce, isPowerOfFourLogarithmic, isPowerOfFourBitwise}
)

func TestPowerOfFour(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			test := test
			t.Run(fmt.Sprint(test.n, test.expected), func(t *testing.T) {
				if have := fn(test.n); have != test.expected {
					t.Errorf(`input: %+vexpected: %+v`, test.expected, have)
				}
			})
		}
	}
}
