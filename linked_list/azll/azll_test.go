package azll

import (
	"fmt"
	"testing"
)

type TestCase struct {
	a any
	b any
	c bool
}

func comparator(tc TestCase) (string, error) {
	cmp := IfEqualAny(tc.a, tc.b)
	msg := fmt.Sprintf("OK: IfEqualAny(%v, %v) = %v, want %v\n", tc.a, tc.b, cmp, tc.c)
	err_msg := fmt.Errorf("Error: IfEqualAny(%v, %v) = %v, want %v\n", tc.a, tc.b, cmp, tc.c)
	if (cmp != tc.c && !tc.c) || (cmp == tc.c && tc.c) || (cmp == tc.c && !tc.c) {
		return msg, nil
	} else {
		return msg, err_msg
	}
}

func TestIfEqualAny(t *testing.T) {

	// Test cases: when both values are valid and equal
	testing_cases := []TestCase{
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