package delete_nodes

import (
	"fmt"
	. "github.com/ashtanko/go-algorithms/ds/linked_list"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestDeleteNodes(t *testing.T) {
	testCases := []struct {
		head     []int
		m        int
		n        int
		expected []int
	}{
		{
			head:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			m:        2,
			n:        3,
			expected: []int{1, 2, 6, 7, 11, 12},
		},
		{
			head:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			m:        1,
			n:        3,
			expected: []int{1, 5, 9},
		},
		{
			head:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			m:        3,
			n:        1,
			expected: []int{1, 2, 3, 5, 6, 7, 9, 10, 11},
		},
		{
			head:     []int{9, 3, 7, 7, 9, 10, 8, 2},
			m:        1,
			n:        2,
			expected: []int{9, 7, 8},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.head)
		t.Run(name, func(t *testing.T) {
			actual := deleteNodes(FromSlice(tc.head), tc.m, tc.n).ToSlice()
			utils.Checkf(t, is.DeepEqual(actual, tc.expected), tc)
		})
	}
}
