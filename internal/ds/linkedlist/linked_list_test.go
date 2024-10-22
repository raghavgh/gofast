package linkedlist

import (
	"testing"
)

// Initial check to see if LinkedList is initialized and both Head and Tail are nil
func Test_Initialize(t *testing.T) {
	list := New()
	if list.Head != nil || list.Tail != nil || list.Len() != 0 {
		t.Error("Linked List not initialized correctly")
	}
}

// Test single and multiple pushes to the front of the list
func Test_PushFront(t *testing.T) {
	list := New()

	list.PushFront(5)
	if list.Head.Val != 5 || list.Tail.Val != 5 || list.Len() != 1 {
		t.Error("Value not correctly pushed to front of linked list")
	}

	list.PushFront(10)
	if list.Head.Val != 10 || list.Tail.Val != 5 || list.Len() != 2 {
		t.Error("Second value not correctly pushed to front of linked list")
	}

	list.PushFront(15)
	if list.Head.Val != 15 || list.Tail.Val != 5 || list.Len() != 3 {
		t.Error("Third value not correctly pushed to front of linked list")
	}
}

// Continuation in next response due to length limitations.

// Test single and multiple pushes to the back of the list
func Test_PushBack(t *testing.T) {
	list := New()

	list.PushBack(5)
	if list.Head.Val != 5 || list.Tail.Val != 5 || list.Len() != 1 {
		t.Error("Value not correctly pushed to back of linked list")
	}

	list.PushBack(10)
	if list.Head.Val != 5 || list.Tail.Val != 10 || list.Len() != 2 {
		t.Error("Second value not correctly pushed to back of linked list")
	}

	list.PushBack(15)
	if list.Head.Val != 5 || list.Tail.Val != 15 || list.Len() != 3 {
		t.Error("Third value not correctly pushed to back of linked list")
	}
}

func Test_MoveToFront(t *testing.T) {
	list := createLinkedList([]int{1, 2, 3, 4, 5})

	// Move middle node to front
	middleNode := list.Head.Next.Next // Value should be 3
	list.MoveToFront(middleNode)
	if list.Head.Val != 3 || list.Len() != 5 {
		t.Errorf("Middle node not correctly moved to front of linked list")
	}

	// Move tail node to front
	tailNode := list.Tail // Value should be 5
	list.MoveToFront(tailNode)
	if list.Head.Val != 5 || list.Len() != 5 {
		t.Errorf("Tail node not correctly moved to front of linked list")
	}
}

func Test_MoveToBack(t *testing.T) {
	list := createLinkedList([]int{1, 2, 3, 4, 5})

	// Move middle node to back
	middleNode := list.Head.Next.Next // Value should be 3
	list.MoveToBack(middleNode)
	if list.Tail.Val != 3 || list.Len() != 5 {
		t.Errorf("Middle node not correctly moved to back of linked list")
	}

	// Move head node to back
	headNode := list.Head // Value should be 1
	list.MoveToBack(headNode)
	if list.Tail.Val != 1 || list.Len() != 5 {
		t.Errorf("Head node not correctly moved to back of linked list")
	}
}

func Test_InsertAfter(t *testing.T) {
	list := createLinkedList([]int{1, 2, 4, 5})

	node := list.Head.Next // Node with value 2
	list.InsertAfter(node, 3)

	// Check if value 3 is properly inserted after value 2
	if node.Next.Val != 3 || list.Len() != 5 {
		t.Error("Value not correctly inserted after the given node")
	}
}

func Test_InsertBefore(t *testing.T) {
	list := createLinkedList([]int{1, 2, 4, 5})

	node := list.Head.Next.Next // Node with value 4
	list.InsertBefore(node, 3)

	// Check if value 3 is properly inserted before value 4
	if node.Prev.Val != 3 || list.Len() != 5 {
		t.Error("Value not correctly inserted before the given node")
	}
}

// Continuation in next response due to length limitations.

// Test removal of nodes
func Test_Remove(t *testing.T) {
	list := createLinkedList([]int{1, 2, 3, 4, 5})

	// Remove middle node
	middleNode := list.Head.Next.Next // Node with value 3
	list.Remove(middleNode)
	if list.Len() != 4 || list.Head.Next.Next.Val != 4 {
		t.Error("Middle node not correctly removed from linked list")
	}

	// Remove head
	headNode := list.Head
	list.Remove(headNode)
	if list.Len() != 3 || list.Head.Val != 2 {
		t.Error("Head node not correctly removed from linked list")
	}

	// Remove tail
	tailNode := list.Tail
	list.Remove(tailNode)
	if list.Len() != 2 || list.Tail.Val != 4 {
		t.Error("Tail node not correctly removed from linked list")
	}

	// Remove from empty list
	emptyList := New()
	emptyNode := &Node{Val: 1}
	emptyList.Remove(emptyNode) // should not panic
}

// Utility function for creating a linked list quickly
func createLinkedList(nums []int) *LinkedList {
	list := New() // Using the New() function defined in your provided code
	for _, num := range nums {
		list.PushBack(num)
	}
	return list
}

// Test removal from empty list
func Test_RemoveFromEmpty(t *testing.T) {
	list := New()
	list.Remove(&Node{Val: 1, Next: nil, Prev: nil})
	// The code should not crash and the list should still be empty.
	if list.Head != nil || list.Tail != nil || list.Len() != 0 {
		t.Error("Remove operation from empty list not handled correctly")
	}
}

// Test PushBack, PushFront, InsertAfter, and InsertBefore on empty list
func Test_PushAndInsertOnEmpty(t *testing.T) {
	list := New()

	// The node becomes the Head and Tail of list simultaneously.
	list.PushBack(1)
	if list.Head.Val != 1 || list.Tail.Val != 1 || list.Len() != 1 {
		t.Error("Push back operation on empty list not handled correctly")
	}

	list = New()
	list.PushFront(1)
	if list.Head.Val != 1 || list.Tail.Val != 1 || list.Len() != 1 {
		t.Error("Push front operation on empty list not handled correctly")
	}

	list = New()
	list.InsertAfter(list.Head, 1)
	if list.Head.Val != 1 || list.Tail.Val != 1 || list.Len() != 1 {
		t.Error("Insert after operation on empty list not handled correctly")
	}

	list = New()
	list.InsertBefore(list.Head, 1)
	if list.Head.Val != 1 || list.Tail.Val != 1 || list.Len() != 1 {
		t.Error("Insert before operation on empty list not handled correctly")
	}
}

// Test MoveToFront and MoveToBack on empty list
func Test_MoveToFrontBackOnEmpty(t *testing.T) {
	list := New()

	// The code should not crash and the list should still be empty.
	list.MoveToFront(&Node{Val: 1, Next: nil, Prev: nil})
	if list.Head != nil || list.Tail != nil || list.Len() != 0 {
		t.Error("Move to front operation on empty list not handled correctly")
	}

	list = New()
	list.MoveToBack(&Node{Val: 1, Next: nil, Prev: nil})
	if list.Head != nil || list.Tail != nil || list.Len() != 0 {
		t.Error("Move to back operation on empty list not handled correctly")
	}
}

// Add more edge test cases as needed.

func Test_InsertBeforeHead(t *testing.T) {
	list := createLinkedList([]int{1, 2, 3, 4, 5})

	// Insert before the head
	beforeNode := list.Head
	list.InsertBefore(beforeNode, 6)

	// Check if value 6 is properly inserted before the old head
	if list.Head.Val != 6 || list.Len() != 6 {
		t.Error("Value not correctly inserted before the head of the list")
	}

	if list.Head.Next.Val != 1 {
		t.Error("Head Node's next pointer not pointing to the previous Head Node")
	}
}

func Test_InsertAfterTail(t *testing.T) {
	list := createLinkedList([]int{1, 2, 3, 4, 5})

	// Insert after the tail
	afterNode := list.Tail
	list.InsertAfter(afterNode, 6)

	// Check if value 6 is properly inserted after the old tail
	if list.Tail.Val != 6 || list.Len() != 6 {
		t.Error("Value not correctly inserted after the tail of the list")
	}

	// Also check if old tail's next node is the new tail node
	if list.Tail.Prev.Val != 5 || afterNode.Next.Val != 6 {
		t.Error("Old Tail Node's next pointer not pointing to the new Tail Node")
	}
}

// Test moving tail to back of list
func Test_MoveToBackTail(t *testing.T) {
	list := createLinkedList([]int{1, 2, 3, 4, 5})

	// Move tail node to back of list
	tailNode := list.Tail
	list.MoveToBack(tailNode)

	// Check if the tail node remains the same and list length remains the same
	if list.Tail.Val != 5 || list.Len() != 5 {
		t.Error("Tail node not correctly handled when moved to back of list")
	}

	// Additional check to see if the tail node is still correctly linked
	if list.Tail.Prev.Val != 4 {
		t.Error("Tail Node's previous node is not correctly linked")
	}
}

// Test removing node when list length is 1
func Test_RemoveSingleElement(t *testing.T) {
	list := createLinkedList([]int{1})

	// Remove sole node in list
	soleNode := list.Head
	list.Remove(soleNode)

	// Check if the head and tail are nil and list length is 0
	if list.Head != nil || list.Tail != nil || list.Len() != 0 {
		t.Error("Removal of sole node in list not handled correctly")
	}
}

// Test moving the head to the front of the list
func Test_MoveToFrontHead(t *testing.T) {
	list := createLinkedList([]int{1, 2, 3, 4, 5})

	// Move head node to front of list
	headNode := list.Head
	list.MoveToFront(headNode)

	// Check if the head node remains the same and list length remains the same
	if list.Head.Val != 1 || list.Len() != 5 {
		t.Error("Head node not correctly handled when moved to front of list")
	}

	// Additional check to see if the head node is still correctly linked
	if list.Head.Next.Val != 2 {
		t.Error("Head Node's next node is not correctly linked")
	}
}
