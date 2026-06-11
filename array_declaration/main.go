package main

import "fmt"

func main() {
    var s1 []string         // nil
    s2 := []string{}        // empty
    s3 := make([]string, 0) // empty, equivalent to s2 := []string{}

    fmt.Printf("s1 is nil: %t, len: %d, cap: %d\n", s1 == nil, len(s1), cap(s1))
    fmt.Printf("s2 is nil: %t, len: %d, cap: %d\n", s2 == nil, len(s2), cap(s2))
    fmt.Printf("s3 is nil: %t, len: %d, cap: %d\n", s3 == nil, len(s3), cap(s3))
}
