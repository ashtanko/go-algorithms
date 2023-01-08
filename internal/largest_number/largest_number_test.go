package largest_number

import (
	"fmt"
	"testing"
)

type strategy struct {
	task func([]int) string
	name string
}

var (
	tests = []struct {
		nums     []int
		expected string
	}{
		{[]int{10, 2}, "210"},
		{[]int{0}, "0"},
		{[]int{3, 30, 34, 5, 9}, "9534330"},
	}

	task = []strategy{{largestNumber, "strategy1"}, {largestNumber2, "strategy2"}}
)

func TestLargestNumber(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			testCaseName := fmt.Sprintf("%s input: %d expected: %s", fn.name, test.nums, test.expected)
			t.Run(testCaseName, func(t *testing.T) {
				t.Parallel()
				if have := fn.task(test.nums); have != test.expected {
					t.Errorf(`input: %+vexpected: %+v`, test.expected, have)
				}
			})
		}
	}
}
