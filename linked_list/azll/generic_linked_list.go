package azll

import (
	"fmt"
	"reflect"
)

// Generic node
type GenericNode[T any] struct {
	Data T
	Next *GenericNode[T]
}

// Generic linked list
type GenericLinkedList[T any] struct {
	Start  *GenericNode[T]
	Length int
}

// Print into the generic linked list
func (l *GenericLinkedList[T]) GenericPrint() {
	temp := l.Start
	for temp != nil {
		fmt.Print(temp.Data, " ")
		temp = temp.Next
	}
	fmt.Println()
}

// Add data to the beginning of the generic linked list
func (l *GenericLinkedList[T]) GenericPrepend(data T) {
	n := GenericNode[T]{
		Data: data,
		Next: nil,
	}
	if l.Start == nil { // If the linked list is empty
		l.Start = &n
		return
	}
	if l.Start.Next == nil { // If it`s` the last node of the linked list
		l.Start.Next = &n
		return
	}
	temp := l.Start
	l.Start = l.Start.Next
	l.GenericPrepend(data)
	l.Start = temp
}

// Add data to the end of the generic linked list
func (l *GenericLinkedList[T]) GenericPostpend(data T) {
	n := GenericNode[T]{
		Data: data,
		Next: nil,
	}
	if l.Start == nil {
		l.Start = &n
		return
	}
	temp := l.Start
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &n
}

// Delete data from the generic linked list
func (l *GenericLinkedList[T]) GenericDeleteNode(value T) error {
	if l.Length == 0 {
		return fmt.Errorf("The linked list is empty")
	}

	// Delete the first node
	if IfEqualAny(l.Start.Data, value) {
		l.Start = l.Start.Next
		l.Length--
		if l.Length == 0 {
			l.Start = nil
		}
		return nil
	}

	prev := l.Start
	for prev.Next != nil || !IfEqualAny(prev.Next.Data, value) {
		prev.Next = prev.Next.Next
		return nil
	}
	prev = prev.Next
	l.Length--
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
