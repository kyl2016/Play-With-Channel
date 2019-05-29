package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type BigData struct {
	data [1024]byte
	eigen [512]byte
	size int
	name string
}

func main() {
	p := sync.Pool{New: func() interface{} {
		return BigData{}
	}}
	fmt.Println(p.Get())

	p.Put(BigData{name:"a"})
	fmt.Println(p.Get())

	p.Put(BigData{name:"b"})
	p.Put(BigData{name:"c"})
	p.Put(BigData{name:"d"})

	fmt.Println(p.Get())
	fmt.Println(p.Get())

	runtime.GC()

	fmt.Println(p.Get())
	fmt.Println(p.New())

	go func() {
		p.Put(BigData{name:"e"})
		p.Put(BigData{name:"f"})
		p.Put(BigData{name:"g"})
		fmt.Println("goroutine 1:", p.Get())
	}()

	go func(){
		time.Sleep(time.Millisecond)
		p.Put(BigData{name:"h"})
		fmt.Println("goroutine 2:", p.Get())
	}()


	go func(){
		time.Sleep(time.Millisecond * 2)
		fmt.Println("goroutine 3:", p.Get())
	}()

	select {

	}
}
