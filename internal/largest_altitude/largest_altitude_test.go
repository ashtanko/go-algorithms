package largest_altitude

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestLargestAltitude(t *testing.T) {
	testCases := []struct {
		gain     []int
		expected int
	}{
		{
			[]int{-5, 1, 5, 0. - 7},
			1,
		},
		{
			[]int{-4, -3, -2, -1, 4, 3, 2},
			0,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.gain)
		t.Run(name, func(t *testing.T) {
			actual := largestAltitude(tc.gain)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
