package longest_consecutive_sequence

import (
	"fmt"
	"testing"
)

var (
	tests = []struct {
		nums     []int
		expected int
	}{
		{
			[]int{100, 4, 200, 1, 3, 2},
			4,
		},
		{
			[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			9,
		},
	}

	task = []func(nums []int) int{longestConsecutive}
)

func TestLongestConsecutiveSequence(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			t.Run(fmt.Sprint(test.nums, test.expected), func(t *testing.T) {
				if have := fn(test.nums); have != test.expected {
					t.Errorf(`input: %+vexpected: %+v`, test.expected, have)
				}
			})
		}
	}
}
