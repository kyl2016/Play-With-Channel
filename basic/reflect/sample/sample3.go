package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.23
	fmt.Println(num)

	pointer := reflect.ValueOf(&num)

	// Elem returns the value that the interface v contains
	// or that the pointer v points to.
	newValue := pointer.Elem()
	newValue.SetFloat(2.34)
	fmt.Printf("%f\n", num)

	// panic
	// pointer := reflect.ValueOf(&num)
	// newValue := pointer.Elem()
	// pointer.CanSet() 返回false
}
