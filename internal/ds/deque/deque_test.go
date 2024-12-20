package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
	assert2 "gotest.tools/v3/assert"
)

func TestNewDeque(t *testing.T) {
	deque := NewDeque()
	assert.NotNil(t, deque)
}

func TestDeque_IsEmpty(t *testing.T) {
	deque := NewDeque()
	assert.True(t, deque.IsEmpty())
	deque.Push(1)
	assert.False(t, deque.IsEmpty())
}

func TestDeque_Inject(t *testing.T) {
	deque := NewDeque()
	deque.Inject(3)
	deque.Inject(2)
	deque.Inject(1)
	expected := []interface{}{3, 2, 1}
	assert2.DeepEqual(t, expected, deque.Items)
}

func TestDeque_Push(t *testing.T) {
	deque := NewDeque()
	deque.Push(3)
	deque.Push(2)
	deque.Push(1)
	expected := []interface{}{1, 2, 3}
	assert2.DeepEqual(t, expected, deque.Items)
}

func TestDeque_Pop(t *testing.T) {
	deque := NewDeque()
	deque.Push(3)
	deque.Push(2)
	deque.Push(1)
	assert.Equal(t, 1, deque.Pop())
	expected := []interface{}{2, 3}
	assert2.DeepEqual(t, expected, deque.Items)
}

func TestDeque_Eject(t *testing.T) {
	deque := NewDeque()
	deque.Push(3)
	deque.Push(2)
	deque.Push(1)
	assert.Equal(t, 3, deque.Eject())
	expected := []interface{}{1, 2}
	assert2.DeepEqual(t, expected, deque.Items)
}

func TestDeque_Size(t *testing.T) {
	deque := NewDeque()
	deque.Push(3)
	deque.Push(2)
	deque.Push(1)
	assert.Equal(t, 3, deque.Size())
}
