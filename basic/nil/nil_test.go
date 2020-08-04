package main

import (
	"fmt"
	"unsafe"
)

// predecalred nil has not a default value
func ExampleNil_DefaultType() {
	// There must be sufficient information for compiler to deduce the type of a nil value.
	_ = (*struct{})(nil)
	_ = []int(nil)
	_ = map[int]bool(nil)
	_ = chan string(nil)
	_ = (func())(nil)
	_ = interface{}(nil)

	// These lines are equivalent to the above lines.
	var _ *struct{} = nil
	var _ []int = nil
	var _ map[int]bool = nil
	var _ chan string = nil
	var _ func() = nil
	var _ interface{} = nil

	// This following line doesn't compile.
	var _ = nil
}

// the sizes of nil values with types of different kinds may be different
func ExampleNil_Size() {
	var p *struct{} = nil
	fmt.Println(unsafe.Sizeof(p)) // 8

	var s []int = nil
	fmt.Println(unsafe.Sizeof(s)) // 24

	var m map[int]bool = nil
	fmt.Println(unsafe.Sizeof(m)) // 8

	var c chan string = nil
	fmt.Println(unsafe.Sizeof(c)) // 8

	var f func() = nil
	fmt.Println(unsafe.Sizeof(f)) // 8

	var i interface{} = nil
	fmt.Println(unsafe.Sizeof(i)) // 16
}

// it is legal to range over nil channels, maps, slices and array pointers
func ExampleNil_Range() {
	// The number of loop steps by iterate a nil array pointer is the length of its corresponding array type.
	for range []int(nil) {
		fmt.Println("Hello")
	}

	for range map[string]string(nil) {
		fmt.Println("world")
	}

	for i := range (*[5]int)(nil) {
		fmt.Println(i)
	}
	// 0
	// 1
	// 2
	// 3
	// 4

	for range chan bool(nil) { // block here
		fmt.Println("Bye")
	}
}

type Slice []bool

func (s Slice) Length() int {
	return len(s)
}

func (s Slice) Modify(i int, x bool) {
	s[i] = x // panic if s is nil
}

func (p *Slice) DoNothing() {
}

func (p *Slice) Append(x bool) {
	*p = append(*p, x) // panic if p is nil
}

// should check p is nil before use it
func (p *Slice) AppendGood(x bool) {
	if p == nil {
		*p = Slice{x}
		return
	}

	*p = append(*p, x)
}

// invoking methods through non-interface nil receiver arguments will not panic
func ExampleNil_ReceiverArguments() {
	// The following selectors will not cause panics.
	_ = ((Slice)(nil)).Length
	_ = ((Slice)(nil)).Modify
	_ = ((*Slice)(nil)).DoNothing
	_ = ((*Slice)(nil)).Append

	// will also not panic
	_ = ((Slice)(nil)).Length()
	((*Slice)(nil)).DoNothing()

	// will panic
	// But panics will not be triggered at the time of invoking the methods.
	// They will be triggered on dereferencing nil pointers in the method bodies.
	((Slice)(nil)).Modify(0, true)
	((*Slice)(nil)).Append(true)

	// output:
}

// FIXME: ?
// *new(T) results a nil T value if the zero value of type T is represented with the predeclared nil identifier
func ExampleNil_NewT() {
	fmt.Println(*new(*int) == nil)
	fmt.Println(*new([]int) == nil)
	fmt.Println(*new(map[int]int) == nil)
	fmt.Println(*new(chan interface{}) == nil)
	fmt.Println(*new(func()) == nil)
	fmt.Println(*new(interface{}) == nil)

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
}
