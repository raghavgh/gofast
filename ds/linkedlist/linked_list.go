package linkedlist

// Node represents a node in a doubly linked list.
type Node struct {
	Val  any
	Next *Node
	Prev *Node
}

// LinkedList represents a doubly linked list.
type LinkedList struct {
	Head *Node
	Tail *Node
	len  int
}

// New returns an initialized list.
func New() *LinkedList {
	return &LinkedList{}
}

// newNode returns a new node with given value
func newNode(val any) *Node {
	return &Node{Val: val}
}

// Len returns the number of elements of list l.
func (l *LinkedList) Len() int {
	return l.len
}

// PushFront inserts a new node with given val at the front of the list
func (l *LinkedList) PushFront(val any) {
	node := newNode(val)
	if l.len == 0 {
		l.Head, l.Tail = node, node
	} else {
		node.Next, l.Head.Prev, l.Head = l.Head, node, node
	}
	l.len++
}

// PushBack inserts a new node with given val at the back of the list
func (l *LinkedList) PushBack(val any) {
	node := newNode(val)
	if l.len == 0 {
		l.Head, l.Tail = node, node
	} else {
		l.Tail.Next, node.Prev, l.Tail = node, l.Tail, node
	}
	l.len++
}

// MoveToFront moves node to the front of the list
func (l *LinkedList) MoveToFront(node *Node) {
	if l.len == 0 || node == l.Head {
		return
	}

	// If the node is at tail, update the tail to the previous node.
	if node == l.Tail {
		l.Tail = node.Prev
	}

	// unlink the node from its current position in the list.
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	// Set the node as new head.
	node.Next, l.Head.Prev = l.Head, node
	node.Prev, l.Head = nil, node
}

// MoveToBack moves node to the back of the list.
func (l *LinkedList) MoveToBack(node *Node) {
	// if the list is empty or the node is already at the tail, no need to move.
	if l.len == 0 || node == l.Tail {
		return
	}

	// If the node is at head, update the head to the next node.
	if node == l.Head {
		l.Head = node.Next
	}

	// unlink the node from its current position in the list.
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	// Set the node as new tail.
	node.Prev, l.Tail.Next = l.Tail, node
	node.Next, l.Tail = nil, node
}

// Remove removes a node from the list
func (l *LinkedList) Remove(node *Node) {
	// If the list is empty, nothing to remove
	if l.len == 0 {
		return
	}

	// Unlink the node from the list
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		// If the node to remove is at the head
		l.Head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		// If the node to remove is at the tail
		l.Tail = node.Prev
	}

	// Clear the value and links on the removed node
	node.Val, node.Prev, node.Next = nil, nil, nil

	// Decrement the length of the list
	l.len--
}

// InsertAfter adds a new node with the given value after a provided node
func (l *LinkedList) InsertAfter(after *Node, val any) {
	// Create new node
	node := newNode(val)

	// When list is empty, make new node the head and tail
	if l.len == 0 {
		l.Head, l.Tail = node, node
	} else {
		// Link node to its next and previous nodes
		node.Prev, node.Next = after, after.Next

		// If inserting after the tail, update tail
		if after == l.Tail {
			l.Tail = node
		} else {
			// Else, update 'Prev' link on node's next node
			after.Next.Prev = node
		}
		// Update 'Next' link on 'after' node
		after.Next = node
	}
	// Increase length of list
	l.len++
}

// InsertBefore adds a new node with the given value before a provided node
func (l *LinkedList) InsertBefore(before *Node, val any) {
	// Create new node
	node := newNode(val)

	// When list is empty, make new node the head and tail
	if l.len == 0 {
		l.Head, l.Tail = node, node
	} else {
		// Link node to its next and previous nodes
		node.Prev, node.Next = before.Prev, before

		// If inserting before the head, update head
		if before == l.Head {
			l.Head = node
		} else {
			// Else, update 'Next' link on node's previous node
			before.Prev.Next = node
		}
		// Update 'Prev' link on 'before' node
		before.Prev = node
	}
	// Increase length of list
	l.len++
}

// Clear removes all nodes from the list.
func (l *LinkedList) Clear() {
	l.Head, l.Tail, l.len = nil, nil, 0
}
