package swap_nodes

import (
	"github.com/ashtanko/go-algorithms/ds/linked_list"
)

func swapNodes(head *linked_list.ListNode, k int) *linked_list.ListNode {
	listLength := 0
	var frontNode *linked_list.ListNode
	var endNode *linked_list.ListNode
	var currentNode = head

	for currentNode != nil {
		listLength++
		if endNode != nil {
			endNode = endNode.Next
		}
		if listLength == k {
			frontNode = currentNode
			endNode = head
		}
		currentNode = currentNode.Next
	}

	if frontNode == nil {
		return head
	}

	if endNode == nil {
		return head
	}

	temp := frontNode.Val
	frontNode.Val = endNode.Val
	endNode.Val = temp
	return head
}
