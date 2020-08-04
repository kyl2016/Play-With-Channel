package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var linkErr *os.LinkError
	tt := reflect.TypeOf(linkErr).Elem()
	fmt.Printf("%+v\n", tt)

	fmt.Printf("type:%v\n", reflect.TypeOf(linkErr))
	fmt.Printf("value:%v\n", reflect.ValueOf(linkErr))

	tty, _ := os.OpenFile("/home/lynxi/tmp/file1.txt", os.O_RDWR, 0)

	var r io.Reader
	r = tty
	fmt.Printf("type:%v\n", reflect.TypeOf(r))

	var w io.Writer
	w = r.(io.Writer)
	fmt.Printf("type:%v\n", reflect.TypeOf(w))

	w = nil
	fmt.Printf("type:%v\n", reflect.TypeOf(w))
	fmt.Printf("value:%v\n", reflect.ValueOf(w))

	// 强制转换
	var num float64 = 1.25
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	convertPointer := pointer.Interface().(*float64)

	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)

	// with ok
	_, ok := pointer.Interface().(float64)
	if !ok {
		panic("can't convert")
	}
}
