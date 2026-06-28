package main

import (
	"fmt"
	"slices"

	"github.com/Azrubael/260601-golang-tutor/llinked_list/azll"
)

func main() {

	// Testing intLinkedList methods
	///////////////////////////////////////////////////////////////////
	mylist := azll.IntLinkedList{}
	intArray := []int{37, 73, 70, 324, 43, 18}

	for _, element := range intArray {
		value := &azll.IntNode{Data: element}
		mylist.Prepend(value)
	}
	fmt.Print(mylist, "\n\n")

	fmt.Println("The created linked list:")
	mylist.PrintListData()

	valToDel := 37
	message := ""
	if slices.Contains(intArray, valToDel) {
		mylist.DeleteWithValue(valToDel)
		message = fmt.Sprintf("The value %d has been deleted from myArray.", valToDel)
	} else {
		message = fmt.Sprintf("The value %d was not found in myArray.", valToDel)
	}
	fmt.Println(message)
	mylist.PrintListData()

	// Insert an element at the end of the list.
	newNode1 := &azll.IntNode{Data: 77}
	mylist.Postpend(newNode1)
	mylist.PrintListData()

	// Insert an element before a given node address
	newNode2 := &azll.IntNode{Data: 55}
	address := mylist.Head.Next
	mylist.InsertBefore(newNode2, address)
	mylist.PrintListData()

	// Testing the deleteWithValue method
	valToDel = 77
	message = ""
	if slices.Contains(intArray, valToDel) {
		mylist.DeleteWithValue(valToDel)
		message = fmt.Sprintf("The value %d has been deleted from myArray.", valToDel)
	} else {
		message = fmt.Sprintf("The value %d was not found in myArray.", valToDel)
	}
	fmt.Println(message)
	mylist.PrintListData()

	// Testing genericLinkedList methods
	///////////////////////////////////////////////////////////////////
	floatLList := azll.GenericLinkedList[float32]{Start: nil, Length: 0}
  floatArray := []float32{37.73, 73.0, 70.5, 324.99, 43.11, 18.001}

  for _, el := range floatArray {
		floatLList.GenericPrepend(el)
	}

  fmt.Println("The created linked list:")
  floatLList.GenericPrint()

  fmt.Println("The updated linked list:")
  floatLList.GenericPostpend(19.19)
  floatLList.GenericPrint()

  stringLList := azll.GenericLinkedList[string]{Start: nil, Length: 0}
  stringArray := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

  for _, el := range stringArray {
		stringLList.GenericPrepend(el)
	}

  fmt.Println("The created linked list:")
  stringLList.GenericPrint()

  fmt.Println("The updated linked list:")
  stringLList.GenericPostpend("eleven")
  stringLList.GenericPrint()

  stringLList.GenericDeleteNode("one")
  fmt.Println("The linked list after deletion:")
  stringLList.GenericPrint()
}