package concurrent_linked_list

import (
	"fmt"
	"strings"
	"sync"
)

type ConcurrentNode struct {
	content interface{}
	next    *ConcurrentNode
}

type ConcurrentLinkedList struct {
	head *ConcurrentNode
	size int
	lock sync.RWMutex
}

func (cll *ConcurrentLinkedList) Append(t interface{}) {
	cll.lock.Lock()
	node := &ConcurrentNode{t, nil}
	if cll.head == nil {
		cll.head = node
	} else {
		last := cll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = node
	}
	cll.size++
	cll.lock.Unlock()
}

func (cll *ConcurrentLinkedList) Insert(i int, t interface{}) error {
	cll.lock.Lock()
	defer cll.lock.Unlock()

	if i < 0 || i > cll.size {
		return fmt.Errorf("Index out of bounds %g", t)
	}
	nodeToAdd := &ConcurrentNode{t, nil}
	if i == 0 {
		nodeToAdd.next = cll.head
		cll.head = nodeToAdd
		return nil
	}
	node := cll.head
	j := 0
	for j < i-2 {
		j++
		node = node.next
	}
	nodeToAdd.next = node.next
	node.next = nodeToAdd
	cll.size++
	return nil
}

func (cll *ConcurrentLinkedList) IsEmpty() bool {
	cll.lock.RLock()
	defer cll.lock.RUnlock()
	if cll.head == nil {
		return true
	}
	return false
}

func (cll *ConcurrentLinkedList) Size() int {
	if cll.head == nil {
		return 0
	}
	cll.lock.RLock()
	defer cll.lock.RLocker()
	size := 1
	last := cll.head

	for {
		if last == nil || last.next == nil {
			break
		}
		last = last.next
		size++
	}

	return size
}

func (cll *ConcurrentLinkedList) String() string {
	cll.lock.RLock()
	defer cll.lock.RUnlock()

	var b strings.Builder
	node := cll.head
	j := 0
	for {
		if node == nil {
			break
		}
		j++
		b.WriteString(fmt.Sprintf("%d", node.content))
		node = node.next
	}
	return b.String()
}

func (cll *ConcurrentLinkedList) Head() *ConcurrentNode {
	cll.lock.RLock()
	defer cll.lock.RUnlock()
	return cll.head
}
