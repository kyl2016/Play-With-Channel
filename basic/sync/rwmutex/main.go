package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	locker := sync.RWMutex{}

	go func() {
		locker.RLock()
		defer locker.RUnlock()

		fmt.Println(time.Now(),"get RLock")

		time.Sleep(time.Second*3)
	}()

	go func() {
		time.Sleep(time.Millisecond*100)
		locker.Lock()
		defer locker.Unlock()

		fmt.Println(time.Now(),"get Lock")

		time.Sleep(time.Second)
	}()

	go func() {
		locker.RLock()
		defer locker.RUnlock()

		fmt.Println(time.Now(),"get RLock 2")
		time.Sleep(time.Second)
	}()

	time.Sleep(time.Second* 5)
}
