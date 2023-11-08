package queue

import "github.com/raghavgh/gofast/ds/linkedlist"

// List wraps a linked list data structure to form a queue
// with optional deletion of arbitrary elements.
//
// If deletion of arbitrary elements is required, the elements added to
// the queue must be distinct or pointers. Failure to adhere to this will
// lead to incorrect behavior because similar elements or non-pointer duplicates
// could interfere with the deletion operation.
//
// The methods Push() and Remove() are designed to handle these scenarios correctly,
// but they rely on the user to ensure the uniqueness of elements when necessary.
type List struct {
	linkedlist.LinkedList
	elementToNode          map[any]*linkedlist.Node
	allowArbitraryDeletion bool
}

// NewQueueList returns an initialized queue.
func NewQueueList(needRemovalFunc bool) *List {
	return &List{
		elementToNode:          make(map[any]*linkedlist.Node),
		allowArbitraryDeletion: needRemovalFunc,
	}
}

// Push inserts a new Element with given val at the back of the queue
func (l *List) Push(element any) {
	l.PushBack(element)
	if l.allowArbitraryDeletion {
		l.elementToNode[element] = l.Tail
	}
}

// Remove removes the given element from the queue
func (l *List) Remove(element any) {
	if node, ok := l.elementToNode[element]; ok {
		l.LinkedList.Remove(node)
		delete(l.elementToNode, element)
	}
}

// Pop removes the first element of the queue
func (l *List) Pop() {
	if l.Empty() {
		panic("queue: Pop() called on empty queue")
	}
	l.Remove(l.Head)
}

// Front returns the first element of the queue
func (l *List) Front() any {
	if l.Empty() {
		panic("queue: Front() called on empty queue")
	}
	return l.Head.Val
}

// Back returns the last element of the queue
func (l *List) Back() any {
	if l.Empty() {
		panic("queue: Back() called on empty queue")
	}
	return l.Tail.Val
}

// Empty returns true if the queue is empty
func (l *List) Empty() bool {
	return l.Len() == 0
}

// Size returns the number of elements of the queue
func (l *List) Size() int {
	return l.Len()
}
