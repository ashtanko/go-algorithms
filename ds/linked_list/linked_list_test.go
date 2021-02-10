package linked_list

import (
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestFromToSlice(t *testing.T) {
	testCases := [][]int{
		{},
		{0, 1, 4, 8},
	}

	for _, testCase := range testCases {
		list := FromSlice(testCase)
		slice := list.ToSlice()

		utils.Checkf(t, is.DeepEqual(slice, testCase), testCase)
	}
}
