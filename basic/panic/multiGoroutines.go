package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				fmt.Println(string(debug.Stack()))
			}
		}()

		panic("err")
	}()

	time.Sleep(time.Millisecond)
}
