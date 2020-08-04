package main

import (
	"fmt"
	"reflect"
)

func main() {
	c := make(chan int)
	m := map[int]string{}
	type People struct {
		Name string
	}
	for _, v := range []interface{}{"hi", 42, func() {}, []int{1, 2}, false, uint8(1), c, m, People{"kitty"}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind:%s, value:%v\n", v.Kind(), v)
		}
	}
}
