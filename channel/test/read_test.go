package test

import "testing"

func TestRead(t *testing.T) {
	println("test")
	ch := make(chan int, 10)
	ch <- 1
	println("begin for")
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				return
			}
			println(msg)
		}
	}
}
