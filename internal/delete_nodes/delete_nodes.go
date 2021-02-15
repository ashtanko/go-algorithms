package delete_nodes

import (
	. "github.com/ashtanko/go-algorithms/ds/linked_list"
)

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}

	currentNode := head

	for {
		for i := 0; i < m-1; i++ {
			if currentNode.Next == nil {
				return head
			}
			currentNode = currentNode.Next
		}

		lastMNode := currentNode
		for i := 0; i < n; i++ {
			if lastMNode.Next == nil {
				break
			}
			lastMNode = lastMNode.Next
		}
		currentNode.Next = lastMNode.Next

		if currentNode.Next == nil {
			return head
		}
		currentNode = currentNode.Next

	}
}
