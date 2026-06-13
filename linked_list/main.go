package main

import (
	"fmt"
	"slices"
)

type node struct{
	data int
	next *node
}

type linkedList struct{
	head *node
	length int
}

func (l *linkedList) prepend (n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData(){
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main(){
	mylist := linkedList{}
	myArray := []int{37, 73, 70, 324, 43, 18}

	for _, element := range myArray {
		value := &node{data: element}
		mylist.prepend(value)
}
	fmt.Print(mylist, "\n\n")

	fmt.Println("The created linked list:")
	mylist.printListData()

	valToDel := 37
	message := ""
	if slices.Contains(myArray, valToDel) {
			mylist.deleteWithValue(valToDel)
			message = fmt.Sprintf("The value %d has been deleted from myArray.", valToDel)
	} else {
		message = fmt.Sprintf("The value %d was not found in myArray.", valToDel)
	}

	fmt.Println(message)
	mylist.printListData()
}