package main

import (
    "fmt"
    "unsafe"
)

type MyStruct struct {
    a int
    b float64
}

type MyInterface interface {
    Method()
}

func main() {
    fmt.Println("=== Basic Types ===")
    fmt.Println("bool:", unsafe.Sizeof(true), "bytes")
    fmt.Println("int:", unsafe.Sizeof(int(0)), "bytes")
    fmt.Println("uint:", unsafe.Sizeof(uint(0)), "bytes")
    fmt.Println("int8:", unsafe.Sizeof(int8(0)), "bytes")
    fmt.Println("uint8:", unsafe.Sizeof(uint8(0)), "bytes")
    fmt.Println("int16:", unsafe.Sizeof(int16(0)), "bytes")
    fmt.Println("uint16:", unsafe.Sizeof(uint16(0)), "bytes")
    fmt.Println("int32:", unsafe.Sizeof(int32(0)), "bytes")
    fmt.Println("uint32:", unsafe.Sizeof(uint32(0)), "bytes")
    fmt.Println("int64:", unsafe.Sizeof(int64(0)), "bytes")
    fmt.Println("uint64:", unsafe.Sizeof(uint64(0)), "bytes")
    fmt.Println("uintptr:", unsafe.Sizeof(uintptr(0)), "bytes")
    fmt.Println("byte:", unsafe.Sizeof(byte(0)), "bytes")
    fmt.Println("rune:", unsafe.Sizeof(rune(0)), "bytes")
    fmt.Println("float32:", unsafe.Sizeof(float32(0)), "bytes")
    fmt.Println("float64:", unsafe.Sizeof(float64(0)), "bytes")
    fmt.Println("complex64:", unsafe.Sizeof(complex64(0)), "bytes")
    fmt.Println("complex128:", unsafe.Sizeof(complex128(0)), "bytes")
    fmt.Println("string:", unsafe.Sizeof(""), "bytes (header only)")

    fmt.Println("\n=== Composite Types ===")
    fmt.Println("array [5]int:", unsafe.Sizeof([5]int{}), "bytes")
    fmt.Println("slice []int:", unsafe.Sizeof([]int{}), "bytes (header only)")
    fmt.Println("map[string]int:", unsafe.Sizeof(map[string]int{}), "bytes (header only)")
    fmt.Println("struct MyStruct:", unsafe.Sizeof(MyStruct{}), "bytes")
    fmt.Println("interface MyInterface:", unsafe.Sizeof((*MyInterface)(nil)), "bytes")
    fmt.Println("function:", unsafe.Sizeof(func() {}), "bytes")
    fmt.Println("channel:", unsafe.Sizeof(make(chan int)), "bytes")
    fmt.Println("pointer *int:", unsafe.Sizeof(new(int)), "bytes")
}
