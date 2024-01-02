package concurrent_linked_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestConcurrentLinkedList_IsEmpty(t *testing.T) {
	t.Run("when empty", func(t *testing.T) {
		list := ConcurrentLinkedList{}
		assert.True(t, list.IsEmpty())
	})

	t.Run("when not empty", func(t *testing.T) {
		list := ConcurrentLinkedList{}
		list.Append(1)
		assert.False(t, list.IsEmpty())
	})
}

func TestConcurrentLinkedList_Size(t *testing.T) {
	t.Run("when empty - should return 0", func(t *testing.T) {
		list := ConcurrentLinkedList{}
		actual := list.Size()
		expected := 0
		assert.Equal(t, expected, actual)
	})

	t.Run("when one element - should return 1", func(t *testing.T) {
		list := ConcurrentLinkedList{}
		list.Append(1)
		actual := list.Size()
		expected := 1
		assert.Equal(t, expected, actual)
	})

	t.Run("when 6 elements - should return 6", func(t *testing.T) {
		list := ConcurrentLinkedList{}
		for i := 0; i < 6; i++ {
			list.Append(i)
		}
		actual := list.Size()
		expected := 6
		assert.Equal(t, expected, actual)
	})
}

func TestConcurrentLinkedList_Append(t *testing.T) {
	t.Run("append one item", func(t *testing.T) {
		list := ConcurrentLinkedList{}
		list.Append(1)
		actual := list.Size()
		assert.Equal(t, 1, actual)
	})

	t.Run("append 10 item", func(t *testing.T) {
		list := ConcurrentLinkedList{}
		for i := 0; i < 10; i++ {
			list.Append(i)
		}
		actual := list.Size()
		assert.Equal(t, 10, actual)
	})

	t.Run("it appends safely concurrently", func(t *testing.T) {
		wantedAppend := 1000
		list := ConcurrentLinkedList{}

		var wg sync.WaitGroup
		wg.Add(wantedAppend)

		for i := 0; i < wantedAppend; i++ {
			go func(w *sync.WaitGroup) {
				list.Append(1)
				w.Done()
			}(&wg)
		}
		wg.Wait()

		actual := list.Size()
		assert.Equal(t, wantedAppend, actual)
	})
}

func TestConcurrentLinkedList_Insert(t *testing.T) {
	list := ConcurrentLinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	err := list.Insert(0, 4)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if size := list.Size(); size != 4 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
	assert.Equal(t, "4123", list.String())
}

func TestConcurrentLinkedList_String(t *testing.T) {
	t.Run("string test", func(t *testing.T) {
		list := ConcurrentLinkedList{}
		list.Append(1)
		list.Append(2)
		actual := list.String()
		assert.Equal(t, "12", actual)
	})
}

func TestConcurrentLinkedList_Head(t *testing.T) {
	list := ConcurrentLinkedList{}
	list.Append("zero")
	h := list.Head()
	if "zero" != fmt.Sprint(h.content) {
		t.Errorf("Expected `zero` but got %s", fmt.Sprint(h.content))
	}
}
