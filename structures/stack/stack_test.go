package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	if stack.items == nil {
		t.Error("stack's inner items should be not nil")
	}
}

func TestStack_Push(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for i := 1; i <= 3; i++ {
		if stack.items[i-1] != i {
			t.Errorf("stack's items: %v, excepted: [1 2 3]", stack.items)
		}
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for i := 1; i <= 3; i++ {
		if stack.items[i-1] != i {
			t.Errorf("stack's items: %v, excepted: [1 2 3]", stack.items)
		}
	}
	top, ok := stack.Pop()
	if !ok || top != 3 {
		t.Errorf("got item: %v, excepted: %v", top, 3)
	}

	stack.Push(3)
	top, ok = stack.Pop()
	if !ok || top != 3 {
		t.Errorf("got item: %v, excepted: %v", top, 3)
	}

	// pop all items
	for i := 0; i < 3; i++ {
		stack.Pop()
	}

	top, ok = stack.Pop()
	if ok || top != nil {
		t.Errorf("got item: %v, excepted: %v", top, 3)
	}
}

func TestStack_Len(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if n := stack.Len(); n != 3 {
		t.Errorf("got len: %v, excepted: %v", n, 3)
	}

	for i := 0; i < 3; i++ {
		stack.Pop()
	}
	if n := stack.Len(); n != 0 {
		t.Errorf("got len: %v, excepted: %v", n, 0)
	}
}
