package main

import (
	"sync"
	"time"
)

var a string
var once sync.Once

func setup() {
	println("setup")
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	println(a)
}

func main() {
	go doprint()
	go doprint()

	time.Sleep(time.Microsecond*10)
}
