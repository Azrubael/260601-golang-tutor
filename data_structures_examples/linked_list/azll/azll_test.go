package azll

import (
	"fmt"
	"testing"
)

type TestCaseEq struct {
	a any
	b any
	c bool
}

func comparator(tc TestCaseEq) (string, error) {
	cmp := IfEqualAny(tc.a, tc.b)
	msg := fmt.Sprintf("OK: IfEqualAny(%v, %v) = %v, want %v\n", tc.a, tc.b, cmp, tc.c)
	err_msg := fmt.Errorf("Error: IfEqualAny(%v, %v) = %v, want %v\n", tc.a, tc.b, cmp, tc.c)
	if (cmp != tc.c && !tc.c) || (cmp == tc.c && tc.c) || (cmp == tc.c && !tc.c) {
		return msg, nil
	} else {
		return msg, err_msg
	}
}

// Test the function 'IfEqualAny(a,b)' with the implementation of safe comparation
func TestIfEqualAny(t *testing.T) {
	testing_cases := []TestCaseEq{
		{1, 1, true},
		{1, 2, false},
		{1.0, 1.0, true},
		{1.0, 1, false},
		{"two", "two", true},
		{"two", "three", false},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{13, 3, 3}, []int{1, 2, 3}, false},
		{[]int{13, 3, 3}, []string{"1", "2", "3"}, false},
		{map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2}, true},
	}

	for i, ts := range testing_cases {
		msg, err := comparator(ts)
		fmt.Println("Test case", i)
		if err != nil {
			t.Errorf(msg, "\n", err)
		} else {
			fmt.Println(msg)
		}
	}

}

func TestGenericDeleteNode(t *testing.T) {
	// Test case 1: Delete the first node
	intArray := GenericLinkedList[int]{}
	intArray.GenericPrepend(1)
	err1 := intArray.GenericDeleteNode(1)
	if err1 != nil {
		t.Errorf("Expected no error, but got %v", err1.Error())
	} else {
		fmt.Printf("The linked list intArray after deleting the only element has lenth: %d\n", intArray.Length)
		intArray.GenericPrint()
	}
	// Test case 2: Delete a node in the middle of the list
	intArray = GenericLinkedList[int]{}
	intArray.GenericPrepend(2)
	intArray.GenericPrepend(1)
	intArray.GenericPrepend(3)
	err2 := intArray.GenericDeleteNode(2)
	if err2 != nil {
				t.Errorf("Expected no error, but got %v", err2.Error())
	} else {
		fmt.Printf("The linked list after deleting the first element %v is: \n", 2)
		intArray.GenericPrint()
	}
	// Test case 3: Delete the last node
	intArray = GenericLinkedList[int]{}
	intArray.GenericPrepend(2)
	intArray.GenericPrepend(1)
	intArray.GenericPrepend(3)
	err3 := intArray.GenericDeleteNode(3)
	if err3 != nil {
				t.Errorf("Expected no error, but got %v", err3.Error())
	} else {
		fmt.Printf("The linked list after deleting the last element %v is: \n", 3)
		intArray.GenericPrint()
	}
}
