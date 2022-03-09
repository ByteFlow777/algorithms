package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stk := NewStack[int]()
	stk.Push(1)
	stk.Pop()
	stk.Top()
	stk.Size()
	stk.Empty()
	stk.typeVal()
}

func TestStack_Push(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for i := 1; i <= 3; i++ {
		if stack.elements[i-1] != i {
			t.Errorf("stack's elements: %v, excepted: [1 2 3]", stack.elements)
		}
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for i := 1; i <= 3; i++ {
		if stack.elements[i-1] != i {
			t.Errorf("stack's elements: %v, excepted: [1 2 3]", stack.elements)
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

	// pop all elements
	for i := 0; i < 3; i++ {
		stack.Pop()
	}

	top, ok = stack.Pop()
	if ok || top != stack.typeVal() {
		t.Errorf("got item: %v, excepted: %v", top, 3)
	}
}

func TestStack_Len(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if n := stack.Size(); n != 3 {
		t.Errorf("got len: %v, excepted: %v", n, 3)
	}

	for i := 0; i < 3; i++ {
		stack.Pop()
	}
	if n := stack.Size(); n != 0 {
		t.Errorf("got len: %v, excepted: %v", n, 0)
	}
}
