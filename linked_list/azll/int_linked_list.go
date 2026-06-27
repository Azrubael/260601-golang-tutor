package azll

import (
	"fmt"
)

// The node of a linked list for integers.
type node struct{
	data int
	next *node
}

// The linked list itself.
type intLinkedList struct{
	head *node
	length int
}

// Prepending into the begin of a linked list.
func (l *intLinkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// Appending into the end of a linked list.
func (l *intLinkedList) postpend(n *node) {
	current := l.head
	if current == nil {
		l.head = n
	} else {
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}
	l.length++
}

// This method takes a pointer to a node (n) and a pointer to the node address before which the new node should be inserted (address).
func (l *intLinkedList) insertBefore(n *node, address *node) {
	if address == nil {
		return
	}
	newNode := &node{data: n.data}
	if address == l.head {
		newNode.next = l.head
		l.head = newNode
	} else {
		current := l.head
		for current != nil && current.next != address {
			current = current.next
		}
		if current != nil {
			newNode.next = address
			current.next = newNode
		}
	}
	l.length++
}

// The method for printing out.
func (l intLinkedList) printListData(){
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

// The method to delete a node with its value.
func (l *intLinkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		if l.length == 0 { l.head = nil }
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}
