package group_anagrams

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	"testing"
)

var (
	tests = []struct {
		strs     []string
		expected [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			[]string{""},
			[][]string{{""}},
		},
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
	}

	task = []func(strs []string) [][]string{groupAnagrams}
)

func TestGroupAnagrams(t *testing.T) {
	for _, fn := range task {
		for _, test := range tests {
			t.Run(fmt.Sprint(test.strs, test.expected), func(t *testing.T) {
				if have := fn(test.strs); !utils.SameString2DSlice(have, test.expected) {
					t.Errorf(`input: %+vexpected: %+v`, test.expected, have)
				}
			})
		}
	}
}
