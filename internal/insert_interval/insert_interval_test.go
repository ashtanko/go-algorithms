package insert_interval

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	tests = []struct {
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
	}

	task = []func(intervals [][]int, newInterval []int) [][]int{insertInterval}
)

func TestInsertInterval(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			test := test
			t.Run(fmt.Sprint(test.intervals, test.expected), func(t *testing.T) {
				if have := fn(test.intervals, test.newInterval); !reflect.DeepEqual(have, test.expected) {
					t.Errorf(`input: %+v expected: %+v`, test.expected, have)
				}
			})
		}
	}
}
