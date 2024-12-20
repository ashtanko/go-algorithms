package middle_node

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/internal/ds/linked_list"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestMiddleNode(t *testing.T) {
	testCases := []struct {
		head     *linked_list.ListNode
		expected []int
	}{
		{
			linked_list.FromSlice([]int{1, 2, 3, 4, 5}),
			[]int{3, 4, 5},
		},
		{
			linked_list.FromSlice([]int{1, 2, 3, 4, 5, 6}),
			[]int{4, 5, 6},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d %d", tc.head.ToSlice(), tc.head.ToSlice())
		t.Run(name, func(t *testing.T) {
			actual := middleNode(tc.head).ToSlice()
			utils.Checkf(t, is.DeepEqual(tc.expected, actual), tc)
		})
	}
}
