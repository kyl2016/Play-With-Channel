package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 0
	fmt.Printf("%T\n", &i)
	fmt.Printf("%T\n", unsafe.Pointer(&i))

	s := "abc"
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", &s)
	fmt.Printf("%T\n", unsafe.Pointer(&s))
}
