package partition_list

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/internal/ds/linked_list"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestPartition(t *testing.T) {
	testCases := []struct {
		head     *linked_list.ListNode
		x        int
		expected *linked_list.ListNode
	}{
		{
			linked_list.FromSlice([]int{1, 4, 3, 2, 5, 2}),
			3,
			linked_list.FromSlice([]int{1, 2, 2, 4, 3, 5}),
		},
		{
			linked_list.FromSlice([]int{2, 1}),
			2,
			linked_list.FromSlice([]int{1, 2}),
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d %d", tc.head.ToSlice(), tc.head.ToSlice())
		t.Run(name, func(t *testing.T) {
			actual := partition(tc.head, tc.x).ToSlice()
			utils.Checkf(t, is.DeepEqual(tc.expected.ToSlice(), actual), tc)
		})
	}
}
