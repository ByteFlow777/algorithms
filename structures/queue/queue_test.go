package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	if queue.items == nil {
		t.Error("queue's inner items should be not nil")
	}
}

func TestQueue_Push(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	for i := 1; i <= 3; i++ {
		if queue.items[i-1] != i {
			t.Errorf("queue's items: %v, excepted: [1 2 3]", queue.items)
		}
	}
}

func TestStack_Pop(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	top, ok := queue.Pop()
	if !ok || top != 1 {
		t.Errorf("got item: %v, excepted: %v", top, 1)
	}

	queue.Push(4)
	top, ok = queue.Pop()
	if !ok || top != 2 {
		t.Errorf("got item: %v, excepted: %v", top, 2)
	}

	// pop all items
	for i := 0; i < 3; i++ {
		queue.Pop()
	}

	top, ok = queue.Pop()
	if ok || top != nil {
		t.Errorf("got item: %v, excepted: %v", top, 3)
	}
}

func TestStack_Len(t *testing.T) {
	queue := NewQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	if n := queue.Len(); n != 3 {
		t.Errorf("got len: %v, excepted: %v", n, 3)
	}

	for i := 0; i < 3; i++ {
		queue.Pop()
	}
	if n := queue.Len(); n != 0 {
		t.Errorf("got len: %v, excepted: %v", n, 0)
	}
}
