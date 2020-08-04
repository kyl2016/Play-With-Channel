package main

import (
	"sync"
	"time"
)
var once sync.Once
var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	println(a)
}

func main() {
	go doprint()
	go doprint()

	time.Sleep(1000)
}
