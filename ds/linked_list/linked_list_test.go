package linked_list

import (
	"github.com/ashtanko/go-algorithms/utils"
	"github.com/stretchr/testify/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestFromToSlice(t *testing.T) {
	testCases := [][]int{
		{},
		{0, 1, 4, 8},
	}

	for _, testCase := range testCases {
		list := FromSlice(testCase)
		slice := list.ToSlice()

		utils.Checkf(t, is.DeepEqual(slice, testCase), testCase)
	}
}

func TestLinkedList_Insert(t *testing.T) {
	t.Run("when empty", func(t *testing.T) {
		list := LinkedList{}
		assert.Equal(t, 0, list.Size())
	})

	t.Run("when full", func(t *testing.T) {
		list := LinkedList{}
		list.Insert(4)
		list.Insert(8)
		list.Insert(15)
		list.Insert(16)
		list.Insert(23)
		list.Insert(42)
		assert.Equal(t, 6, list.Size())
	})
}

func TestLinkedList_Display(t *testing.T) {
	t.Run("when one element", func(t *testing.T) {
		list := LinkedList{}
		list.Insert(1)
		expected := "1 -> "
		actual := list.Display()
		assert.Equal(t, expected, actual)
	})

	t.Run("when two elements", func(t *testing.T) {
		list := LinkedList{}
		list.Insert(1)
		list.Insert(2)
		expected := "2 -> 1 -> "
		actual := list.Display()
		assert.Equal(t, expected, actual)
	})

	t.Run("when 3 elements", func(t *testing.T) {
		list := LinkedList{}
		list.Insert(1)
		list.Insert(2)
		list.Insert(3)
		expected := "3 -> 2 -> 1 -> "
		actual := list.Display()
		assert.Equal(t, expected, actual)
	})
}

func TestLinkedList_ShowBackwards(t *testing.T) {
	t.Run("when one element", func(t *testing.T) {
		list := LinkedList{}
		list.Insert(1)
		expected := "1 <- "
		actual := list.ShowBackwards()
		assert.Equal(t, expected, actual)
	})

	t.Run("when two elements", func(t *testing.T) {
		list := LinkedList{}
		list.Insert(1)
		list.Insert(2)
		expected := "1 <- 2 <- "
		actual := list.ShowBackwards()
		assert.Equal(t, expected, actual)
	})

	t.Run("when 3 elements", func(t *testing.T) {
		list := LinkedList{}
		list.Insert(1)
		list.Insert(2)
		list.Insert(3)
		expected := "1 <- 2 <- 3 <- "
		actual := list.ShowBackwards()
		assert.Equal(t, expected, actual)
	})
}

func TestLinkedList_Reverse(t *testing.T) {
	list := LinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Reverse()
	expected := "1 -> 2 -> 3 -> "
	actual := list.Display()
	assert.Equal(t, expected, actual)
}
