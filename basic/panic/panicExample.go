package main

import "fmt"

func main() {
	fmt.Println("Enter function main.")
	caller1()
	fmt.Println("Enter function main.")
}

func caller1() {
	fmt.Println("Enter function caller1.")
	caller2()
	fmt.Println("Enter function caller1.")
}

func caller2() {
	fmt.Println("Enter function caller2.")
	s := []int{0, 1, 2, 3, 4}
	e5 := s[5]
	_ = e5
	fmt.Println("Exit function caller2")
}

//panic: runtime error: index out of range
//
//goroutine 1 [running]:
//main.caller2()
///home/lynxi/go/src/github.com/kyl2016/Play-With-Golang/panic/panicExample.go:20 +0x62
//main.caller1()
///home/lynxi/go/src/github.com/kyl2016/Play-With-Golang/panic/panicExample.go:13 +0x66
//main.main()
///home/lynxi/go/src/github.com/kyl2016/Play-With-Golang/panic/panicExample.go:7 +0x66
//exit status 2
