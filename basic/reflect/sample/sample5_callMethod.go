package main

import (
	"fmt"
	"reflect"
)

func main() {
	user := User{1, "Kitty", 23}
	getValue := reflect.ValueOf(&user)
	method := getValue.MethodByName("Set")
	args := []reflect.Value{reflect.ValueOf("hello kitty"), reflect.ValueOf(22)}
	method.Call(args)

	fmt.Printf("%+v", user)
}
