package main

import (
	"errors"
	"fmt"
	"reflect"
)

func foo() {
	fmt.Println("call func foo")
}

func bar(a, b, c int) {
	fmt.Println("call bar with ", a, b, c)
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func main() {
	funcs := map[string]interface{}{
		"foo": foo,
		"bar": bar,
	}
	Call(funcs, "foo")
	Call(funcs, "bar", 1, 2, 3)
	Call(funcs, "bar", 1, 2, 3, 4)
}
