package total_money

import (
	"fmt"
	"testing"
)

var (
	tests = []struct {
		n        int
		expected int
	}{
		{4, 10},
		{10, 37},
		{20, 96},
	}

	task = []func(n int) int{totalMoneySimulate, totalMoneyMath}
)

func TestTotalMoney(t *testing.T) {
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
