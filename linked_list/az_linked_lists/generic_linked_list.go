package az_linked_lists

import (
	"fmt"
	"reflect"
)

// Generic node
type genericNode[T any] struct {
	Data T
	next *genericNode[T]
}

// Generic linked list
type genericLinkedList[T any] struct {
	start  *genericNode[T]
	length int
}

// Print into the generic linked list
func (l *genericLinkedList[T]) genericPrint() {
	temp := l.start
	for temp != nil {
		fmt.Print(temp.Data, " ")
		temp = temp.next
	}
	fmt.Println()
}

// Add data to the beginning of the generic linked list
func (l *genericLinkedList[T]) genericPrepend(data T) {
	n := genericNode[T]{
		Data: data,
		next: nil,
	}
	if l.start == nil { // If the linked list is empty
		l.start = &n
		return
	}
	if l.start.next == nil { // If it`s` the last node of the linked list
		l.start.next = &n
		return
	}
	temp := l.start
	l.start = l.start.next
	l.genericPrepend(data)
	l.start = temp
}

// Add data to the end of the generic linked list
func (l *genericLinkedList[T]) genericPostpend(data T) {
	n := genericNode[T]{
		Data: data,
		next: nil,
	}
	if l.start == nil {
		l.start = &n
		return
	}
	temp := l.start
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &n
}

// Delete data from the generic linked list
func (l *genericLinkedList[T]) genericDeleteNode(value T) error {
	if l.length == 0 {
		return fmt.Errorf("The linked list is empty")
	}

	// Delete the first node
	if IfEqualAny(l.start.Data, value) {
		l.start = l.start.next
		l.length--
		if l.length == 0 {
			l.start = nil
		}
		return nil
	}

	prev := l.start
	for prev.next != nil || !IfEqualAny(prev.next.Data, value) {
		prev.next = prev.next.next
		return nil
	}
	prev = prev.next
	l.length--
	return fmt.Errorf("The value %v was not found in the linked list", value)
}

func IfEqualAny[T any](a, b T) bool {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	if !av.IsValid() && !bv.IsValid() {
		return true
	}
	if !av.IsValid() || !bv.IsValid() {
		return false
	}

	kindNilCapable := func(k reflect.Kind) bool {
		switch k {
		case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Func, reflect.Chan:
			return true
		default:
			return false
		}
	}

	if kindNilCapable(av.Kind()) && av.IsNil() {
		return kindNilCapable(bv.Kind()) && bv.IsNil()
	}
	if kindNilCapable(bv.Kind()) && bv.IsNil() {
		return false
	}

	if av.Type() != bv.Type() {
		return false
	}

	return reflect.DeepEqual(av.Interface(), bv.Interface())
}
