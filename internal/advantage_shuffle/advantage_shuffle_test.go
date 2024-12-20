package advantage_shuffle

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestAdvantageShuffle(t *testing.T) {
	testCases := []struct {
		a        []int
		b        []int
		expected []int
	}{
		{
			[]int{2, 7, 11, 15},
			[]int{1, 10, 4, 11},
			[]int{2, 11, 7, 15},
		},
		{
			[]int{12, 24, 8, 32},
			[]int{13, 25, 32, 11},
			[]int{24, 32, 8, 12},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d %d", tc.a, tc.b)
		t.Run(name, func(t *testing.T) {
			actual := advantageCount(tc.a, tc.b)
			fmt.Println("ACTUAL: ", actual)
			utils.Checkf(t, is.DeepEqual(tc.expected, actual), tc)
		})
	}
}
