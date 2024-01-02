package add_two_numbers

import (
	. "github.com/ashtanko/go-algorithms/internal/ds/linked_list"
)

// Function to add two numbers represented as linked lists.
// Each node in the lists contains a single digit.
// The digits are stored in reverse order, and the first node contains the least significant digit.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Initialize a dummy node and a variable to hold the sum of the current digits.
	dummy, sum := new(ListNode), 0
	// Loop through the lists until we reach the end of both lists and there is no carry left.
	for cur := dummy; l1 != nil || l2 != nil || sum != 0; cur = cur.Next {
		// If the first list is not empty, add the current digit to the sum and move to the next node.
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		// If the second list is not empty, add the current digit to the sum and move to the next node.
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		// Create a new node with the last digit of the sum and add it to the result list.
		cur.Next = &ListNode{Val: sum % 10}
		// Remove the last digit from the sum.
		sum /= 10
	}
	// Return the result list, skipping the dummy node.
	return dummy.Next
}
