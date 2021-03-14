package linked_list

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type LinkedList struct {
	head *Node
	tail *Node
}

func FromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	first := &ListNode{Val: nums[0]}
	for current, i := first, 1; i < len(nums); i++ {
		next := ListNode{Val: nums[i]}
		current.Next = &next
		current = &next
	}
	return first
}

func (list *ListNode) ToSlice() []int {
	if list == nil {
		return []int{}
	}

	var slice []int
	for list != nil {
		slice = append(slice, list.Val)
		list = list.Next
	}
	return slice
}

func (L *LinkedList) ToSlice() []interface{} {
	curr := L.head
	if L == nil {
		return []interface{}{}
	}

	var result []interface{}

	for curr != nil {
		result = append(result, curr.key)
		curr = curr.next
	}

	return result
}

func (L *LinkedList) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head

	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (L *LinkedList) Delete(key interface{}) {

}

func (L *LinkedList) Display() string {
	list := L.head
	b := strings.Builder{}
	for list != nil {
		b.WriteString(fmt.Sprintf("%+v -> ", list.key))
		list = list.next
	}
	return b.String()
}

func (L *LinkedList) ShowBackwards() string {
	list := L.tail
	b := strings.Builder{}
	for list != nil {
		b.WriteString(fmt.Sprintf("%v <- ", list.key))
		list = list.prev
	}
	return b.String()
}

func (L *LinkedList) Size() int {
	size := 0
	curr := L.head
	for curr != nil {
		curr = curr.next
		size++
	}
	return size
}

func (L *LinkedList) Reverse() {
	curr := L.head
	var prev *Node
	L.tail = L.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	L.head = prev
}
