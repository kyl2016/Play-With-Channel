package main

import "time"

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second)
		close(ch)
	}()

	ch <- 1
}
