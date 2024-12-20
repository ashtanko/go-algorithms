package swap_nodes

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/internal/ds/linked_list"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

func TestSwapNodes(t *testing.T) {
	testCases := []struct {
		head     *linked_list.ListNode
		k        int
		expected *linked_list.ListNode
	}{
		{
			&linked_list.ListNode{
				Val: 1,
				Next: &linked_list.ListNode{
					Val: 2,
					Next: &linked_list.ListNode{
						Val: 3,
						Next: &linked_list.ListNode{
							Val: 4,
							Next: &linked_list.ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			2,
			&linked_list.ListNode{
				Val: 1,
				Next: &linked_list.ListNode{
					Val: 4,
					Next: &linked_list.ListNode{
						Val: 3,
						Next: &linked_list.ListNode{
							Val: 2,
							Next: &linked_list.ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			&linked_list.ListNode{
				Val: 7,
				Next: &linked_list.ListNode{
					Val: 9,
					Next: &linked_list.ListNode{
						Val: 6,
						Next: &linked_list.ListNode{
							Val: 6,
							Next: &linked_list.ListNode{
								Val: 7,
								Next: &linked_list.ListNode{
									Val: 8,
									Next: &linked_list.ListNode{
										Val: 3,
										Next: &linked_list.ListNode{
											Val: 0,
											Next: &linked_list.ListNode{
												Val: 9,
												Next: &linked_list.ListNode{
													Val: 5,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			5,
			&linked_list.ListNode{
				Val: 7,
				Next: &linked_list.ListNode{
					Val: 9,
					Next: &linked_list.ListNode{
						Val: 6,
						Next: &linked_list.ListNode{
							Val: 6,
							Next: &linked_list.ListNode{
								Val: 8,
								Next: &linked_list.ListNode{
									Val: 7,
									Next: &linked_list.ListNode{
										Val: 3,
										Next: &linked_list.ListNode{
											Val: 0,
											Next: &linked_list.ListNode{
												Val: 9,
												Next: &linked_list.ListNode{
													Val: 5,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			&linked_list.ListNode{
				Val: 1,
			},
			1,
			&linked_list.ListNode{
				Val: 1,
			},
		},
		{
			&linked_list.ListNode{
				Val: 1,
				Next: &linked_list.ListNode{
					Val: 2,
				},
			},
			1,
			&linked_list.ListNode{
				Val: 2,
				Next: &linked_list.ListNode{
					Val: 1,
				},
			},
		},
		{
			&linked_list.ListNode{
				Val: 1,
				Next: &linked_list.ListNode{
					Val: 2,
					Next: &linked_list.ListNode{
						Val: 3,
					},
				},
			},
			2,
			&linked_list.ListNode{
				Val: 1,
				Next: &linked_list.ListNode{
					Val: 2,
					Next: &linked_list.ListNode{
						Val: 3,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d", tc.head.ToSlice())
		t.Run(name, func(t *testing.T) {
			actual := swapNodes(tc.head, tc.k).ToSlice()
			expected := tc.expected.ToSlice()
			fmt.Println(actual, expected)
			utils.Checkf(t, is.DeepEqual(expected, actual), tc)
		})
	}
}
