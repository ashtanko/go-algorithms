package middle_node

import . "github.com/ashtanko/go-algorithms/ds/linked_list"

func middleNode(head *ListNode) *ListNode {

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
