package two_sum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tests = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{}, 0, []int{}},
	}

	task = []func(nums []int, target int) []int{twoSum}
)

func TestTwoSum(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			t.Run(fmt.Sprint(test.nums, test.expected), func(t *testing.T) {
				if have := fn(test.nums, test.target); !assert.Equal(t, have, test.expected) {
					t.Errorf(`input: %+vexpected: %+v`, test.expected, have)
				}
			})
		}
	}
}
