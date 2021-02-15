package swap_pairs

import . "github.com/ashtanko/go-algorithms/ds/linked_list"

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := head.Next
	head.Next = swapPairs(head.Next.Next)
	result.Next = head
	return result
}
