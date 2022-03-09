package stack

import "fmt"

type Stack[T any] struct {
	elements     []T
	elementCount int
}

// NewStack return a new Stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push an element to the top of Stack
func (stk *Stack[T]) Push(element T) {
	stk.elements = append(stk.elements, element)
	stk.elementCount++
}

// typeVal return default val of type T
func (stk *Stack[T]) typeVal() (val T) {
	return val
}

// Pop the top element of Stack and remove it
// If stk.Size() == 0 mean no element in Stack, return default value of type T
func (stk *Stack[T]) Pop() (T, bool) {
	if stk.Size() == 0 {
		return stk.typeVal(), false
	}
	stk.elementCount--
	topElement := stk.elements[stk.elementCount]
	stk.elements = stk.elements[:stk.elementCount]
	return topElement, true
}

// Top return the top element of Stack, but don't remove it
func (stk *Stack[T]) Top() (T, bool) {
	if stk.Size() == 0 {
		return stk.typeVal(), false
	}
	return stk.elements[stk.elementCount], true
}

// Size return the number of Stack's elements
func (stk *Stack[T]) Size() int {
	return stk.elementCount
}

// Empty return whether Stack is nil
// true mean stack is nil, false mean not
func (stk *Stack[T]) Empty() bool {
	return stk.elementCount == 0
}

func (stk *Stack[T]) String() string {
	return fmt.Sprintf("stack%v", stk.elements)
}
