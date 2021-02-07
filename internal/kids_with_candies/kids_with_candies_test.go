package kids_with_candies

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	tests = []struct {
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{[]int{}, 0, []bool{}},
		{[]int{}, 50_000, []bool{}},
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
	}
	task = []func(candies []int, extraCandies int) []bool{kidsWithCandies}
)

func TestRemoveVowels(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			test := test
			t.Run(fmt.Sprint(test.candies, test.extraCandies, test.expected), func(t *testing.T) {
				t.Parallel()
				if have := fn(test.candies, test.extraCandies); !reflect.DeepEqual(have, test.expected) {
					t.Errorf(`input: %+vexpected: %+v`, test.expected, have)
				}
			})
		}
	}
}
