package main

import "fmt"

func main() {
    var s1 []string         // nil
    s2 := []string{}        // empty
    s3 := make([]string, 0) // empty, equivalent to s2 := []string{}

    fmt.Printf("s1 is nil: %t, len: %d, cap: %d\n", s1 == nil, len(s1), cap(s1))
    fmt.Printf("s2 is nil: %t, len: %d, cap: %d\n", s2 == nil, len(s2), cap(s2))
    fmt.Printf("s3 is nil: %t, len: %d, cap: %d\n", s3 == nil, len(s3), cap(s3))
	
	// Create an empty map with string keys and int values using 'make' command
	scores := make(map[string]int)
	scores["Alice"] = 90
	scores["Bob"] = 85

	// Declare and initialize a map in one step with map literal
	fruits := map[string]string{
		"a": "Apple",
		"b": "Banana",
		"c": "Cherry",
	}
	
	// Declare a map variable without initializing (var declaration)
	var ages map[string]int
	ages = make(map[string]int)
	ages["Tom"] = 25
	ages["Jerry"] = 22
	ages["Guffy"] = 17
	
	delete(ages, "Guffy")

}

}
