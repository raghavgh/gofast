package stack

import (
	"github.com/raghavgh/gofast/internal/ds/linkedlist"
)

// Stack wraps a linked list data structure to form a stack
// with optional deletion of arbitrary elements.
//
// If deletion of arbitrary elements is required, the elements added to
// the stack must be distinct or pointers. Failure to adhere to this will
// lead to incorrect behavior because similar elements or non-pointer duplicates
// could interfere with the deletion operation.
//
// The methods Push() and Remove() are designed to handle these scenarios correctly,
// but they rely on the user to ensure the uniqueness of elements when necessary.
type Stack struct {
	linkedlist.LinkedList
	elementToNode          map[any]*linkedlist.Node
	allowArbitraryDeletion bool
}

// NewStack returns an initialized stack.
func NewStack(needRemovalFunc bool) *Stack {
	return &Stack{
		elementToNode:          map[any]*linkedlist.Node{},
		allowArbitraryDeletion: needRemovalFunc,
	}
}

// Push inserts a new Element with given val at the front of the stack
func (s *Stack) Push(element any) {
	s.PushFront(element)
	if s.allowArbitraryDeletion {
		s.elementToNode[element] = s.Head
	}
}

// Remove removes the given element from the stack
func (s *Stack) Remove(element any) {
	if node, ok := s.elementToNode[element]; ok {
		s.LinkedList.Remove(node)
		delete(s.elementToNode, element)
	}
}

// Pop removes the first element of the stack
func (s *Stack) Pop() {
	if s.Empty() {
		panic("stack: Pop() called on empty stack")
	}
	s.Remove(s.Head.Val)
}

// Top returns the first element of the stack
func (s *Stack) Top() any {
	if s.Empty() {
		panic("stack: Top() called on empty stack")
	}
	return s.Head.Val
}

// Empty returns true if the stack is empty
func (s *Stack) Empty() bool {
	return s.Len() == 0
}

// Size returns the number of elements of the stack
func (s *Stack) Size() int {
	return s.Len()
}
