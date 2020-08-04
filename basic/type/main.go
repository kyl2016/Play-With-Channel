package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{} = 1
	fmt.Println("1 type is", reflect.TypeOf(i).Name())
	fmt.Println(reflect.Int)
	fmt.Println(reflect.Int.String())

	printType(map[interface{}]interface{}{})
	printType([]int{})
	printType(func() {})
	printType(func(i int) {})
	printType(make(chan int))

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	fmt.Println(reflect.TypeOf(ch1) == reflect.TypeOf(ch2))

	ch3 := make(chan int)
	fmt.Println(reflect.TypeOf(ch1) == reflect.TypeOf(ch3))

	ch4 := make(chan string)
	printType(ch4)
	fmt.Println(reflect.TypeOf(ch1) == reflect.TypeOf(ch4))
}

func printType(m interface{}) {
	t := reflect.TypeOf(m)
	fmt.Println(t, "Comparable:", t.Comparable())
}
