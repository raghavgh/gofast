package stack

import "github.com/raghavgh/gofast/ds/linkedlist"

// Stack represents a stack.
type Stack struct {
	linkedlist.LinkedList
}

// NewStack returns an initialized stack.
func NewStack() *Stack {
	return &Stack{}
}

// Push inserts a new Element with given val at the front of the stack
func (l *Stack) Push(element any) {
	l.PushFront(element)
}

// Pop removes the first element of the stack
func (l *Stack) Pop() {
	if l.Empty() {
		panic("stack: Pop() called on empty stack")
	}
	l.Remove(l.Head)
}

// Top returns the first element of the stack
func (l *Stack) Top() any {
	if l.Empty() {
		panic("stack: Top() called on empty stack")
	}
	return l.Head.Val
}

// Empty returns true if the stack is empty
func (l *Stack) Empty() bool {
	return l.Len() == 0
}

// Size returns the number of elements of the stack
func (l *Stack) Size() int {
	return l.Len()
}
