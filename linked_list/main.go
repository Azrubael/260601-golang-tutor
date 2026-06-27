package main

import (
	"fmt"
	"slices"
)

func main() {

	// Testing intLinkedList methods
	///////////////////////////////////////////////////////////////////
	mylist := azll.intLinkedList{}
	intArray := []int{37, 73, 70, 324, 43, 18}

	for _, element := range intArray {
		value := &azll.node{data: element}
		mylist.prepend(value)
	}
	fmt.Print(mylist, "\n\n")

	fmt.Println("The created linked list:")
	mylist.printListData()

	valToDel := 37
	message := ""
	if slices.Contains(intArray, valToDel) {
		mylist.deleteWithValue(valToDel)
		message = fmt.Sprintf("The value %d has been deleted from myArray.", valToDel)
	} else {
		message = fmt.Sprintf("The value %d was not found in myArray.", valToDel)
	}
	fmt.Println(message)
	mylist.printListData()

	// Insert an element at the end of the list.
	newNode1 := &azll.node{data: 77}
	mylist.postpend(newNode1)
	mylist.printListData()

	// Insert an element before a given node address
	newNode2 := &azll.node{data: 55}
	address := mylist.head.next
	mylist.insertBefore(newNode2, address)
	mylist.printListData()

	// Testing the deleteWithValue method
	valToDel = 77
	message = ""
	if slices.Contains(intArray, valToDel) {
		mylist.deleteWithValue(valToDel)
		message = fmt.Sprintf("The value %d has been deleted from myArray.", valToDel)
	} else {
		message = fmt.Sprintf("The value %d was not found in myArray.", valToDel)
	}
	fmt.Println(message)
	mylist.printListData()

	// Testing genericLinkedList methods
	///////////////////////////////////////////////////////////////////
	floatLList := azll.genericLinkedList[float32]{start: nil, length: 0}
  floatArray := []float32{37.73, 73.0, 70.5, 324.99, 43.11, 18.001}

  for _, el := range floatArray {
		floatLList.genericPrepend(el)
	}

  fmt.Println("The created linked list:")
  floatLList.genericPrint()

  fmt.Println("The updated linked list:")
  floatLList.genericPostpend(19.19)
  floatLList.genericPrint()

  stringLList := azll.genericLinkedList[string]{start: nil, length: 0}
  stringArray := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

  for _, el := range stringArray {
		stringLList.genericPrepend(el)
	}

  fmt.Println("The created linked list:")
  stringLList.genericPrint()

  fmt.Println("The updated linked list:")
  stringLList.genericPostpend("eleven")
  stringLList.genericPrint()

  stringLList.genericDeleteNode("one")
  fmt.Println("The linked list after deletion:")
  stringLList.genericPrint()
}