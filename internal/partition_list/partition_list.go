package partition_list

import (
	. "github.com/ashtanko/go-algorithms/internal/ds/linked_list"
)

func partition(head *ListNode, x int) *ListNode {
	cur := head
	var l1 []int
	var l2 []int
	for cur != nil {
		v := cur.Val
		if v < x {
			l1 = append(l1, v)
		} else {
			l2 = append(l2, v)
		}
		cur = cur.Next
	}
	l := append(l1, l2...)
	t := 0
	curr := head
	for curr != nil {
		curr.Val = l[t]
		t++
		curr = curr.Next
	}
	return head
}
