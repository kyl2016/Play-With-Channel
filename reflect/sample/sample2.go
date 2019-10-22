package main

import (
	"fmt"
	"reflect"
)

func main() {
	user := User{1, "Kitty", 20}
	reflectFiledsAndMethods(user)
}

func reflectFiledsAndMethods(input interface{}) {
	_type := reflect.TypeOf(input)
	value := reflect.ValueOf(input)

	filedCount := _type.NumField()
	for i := 0; i < filedCount; i++ {
		fmt.Printf("field type : %#v\n", _type.Field(i))
		fmt.Printf("field value: %#v\n", value.Field(i).Interface())
	}

	methodCount := _type.NumMethod()
	for i := 0; i < methodCount; i++ {
		fmt.Printf("method type : %#v\n", _type.Method(i))
		fmt.Printf("method value: %#v\n", value.Method(i).Interface())
	}
}
