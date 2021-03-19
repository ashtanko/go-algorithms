package keys_and_rooms

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestKeysAndRooms(t *testing.T) {
	testCases := []struct {
		rooms    [][]int
		expected bool
	}{
		{
			[][]int{{1}, {2}, {3}, {}},
			true,
		},
		{
			[][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			false,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.rooms)
		t.Run(name, func(t *testing.T) {
			actual := canVisitAllRooms(tc.rooms)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}
