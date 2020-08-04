package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var i = nil // can't assign nil without explicit type

	var err error
	fmt.Println(reflect.TypeOf(err))
	// <nil>

	fmt.Println("begin")
	err = newMyError()
	fmt.Println(reflect.TypeOf(err))
	//  *myError
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("end")
}

type myError struct {
}

func newMyError() *myError {
	return nil
}

func (m *myError) Error() string {
	return ""
}
