package azll

import (
	"fmt"
)

// The node of a linked list for integers.
type IntNode struct{
	Data int
	Next *IntNode
}

// The linked list itself.
type IntLinkedList struct{
	Head *IntNode
	length int
}

// Prepending into the begin of a linked list.
func (l *IntLinkedList) Prepend(n *IntNode) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.length++
}

// Appending into the end of a linked list.
func (l *IntLinkedList) Postpend(n *IntNode) {
	current := l.Head
	if current == nil {
		l.Head = n
	} else {
		for current.Next != nil {
			current = current.Next
		}
		current.Next = n
	}
	l.length++
}

// This method takes a pointer to a node (n) and a pointer to the node address before which the new node should be inserted (address).
func (l *IntLinkedList) InsertBefore(n *IntNode, address *IntNode) {
	if address == nil {
		return
	}
	newNode := &IntNode{Data: n.Data}
	if address == l.Head {
		newNode.Next = l.Head
		l.Head = newNode
	} else {
		current := l.Head
		for current != nil && current.Next != address {
			current = current.Next
		}
		if current != nil {
			newNode.Next = address
			current.Next = newNode
		}
	}
	l.length++
}

// The method for printing out.
func (l IntLinkedList) PrintListData(){
	toPrint := l.Head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
		l.length--
	}
	fmt.Println()
}

// The method to delete a node with its value.
func (l *IntLinkedList) DeleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.Head.Data == value {
		l.Head = l.Head.Next
		l.length--
		if l.length == 0 { l.Head = nil }
		return
	}
	previousToDelete := l.Head
	for previousToDelete.Next.Data != value {
		previousToDelete = previousToDelete.Next
	}
	previousToDelete.Next = previousToDelete.Next.Next
	l.length--
}
