package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

var state string

func test(x int) {
	cond.L.Lock()
	fmt.Println("aaa:", x)
	cond.Wait()
	fmt.Println("bbb:", x)
	//time.Sleep(time.Second * 2)
	cond.L.Unlock()
}

func main() {
	for i := 0; i < 5; i++ {
		go test(i)
	}

	go func() {
		fmt.Println("start all")
		time.Sleep(time.Second * 1)
		fmt.Println("broadcast")
		cond.Signal()
		time.Sleep(time.Second * 1)
		cond.Signal()
		time.Sleep(time.Second * 1)
		//cond.Broadcast()
		//time.Sleep(time.Second * 10)
		fmt.Println("finish all")
	}()

	time.Sleep(time.Minute)
}
