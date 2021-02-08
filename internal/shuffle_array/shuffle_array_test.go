package shuffle_array

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	tests = []struct {
		input    []int
		n        int
		expected []int
	}{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}

	task = []func(nums []int, n int) []int{shuffle}
)

func TestShuffleArray(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			test := test
			t.Run(fmt.Sprint(test.input, test.n), func(t *testing.T) {
				t.Parallel()
				if actual := fn(test.input, test.n); !reflect.DeepEqual(actual, test.expected) {
					t.Errorf(`input: %+vexpected: %+v`, test.expected, actual)
				}
			})
		}
	}
}
