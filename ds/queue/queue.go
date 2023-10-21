package queue

import "github.com/raghavgh/gofast/ds/linkedlist"

// List Queue embeds linkedlist.LinkedList
// created different struct to restrict the usage of queue and functions.
type List struct {
	linkedlist.LinkedList
}

// NewQueueList returns an initialized queue.
func NewQueueList() *List {
	return &List{}
}

// Push inserts a new Element with given val at the back of the queue
func (l *List) Push(element any) {
	l.PushBack(element)
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
