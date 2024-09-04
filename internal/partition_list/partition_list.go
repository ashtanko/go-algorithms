package partition_list

import (
	. "github.com/ashtanko/go-algorithms/internal/ds/linked_list"
)

func partition(head *ListNode, x int) *ListNode {
	current := head
	var lessThanX []int
	var greaterOrEqualX []int
	for current != nil {
		value := current.Val
		if value < x {
			lessThanX = append(lessThanX, value)
		} else {
			greaterOrEqualX = append(greaterOrEqualX, value)
		}
		current = current.Next
	}
	combined := append(lessThanX, greaterOrEqualX...)
	index := 0
	current = head
	for current != nil {
		current.Val = combined[index]
		index++
		current = current.Next
	}
	return head
}
