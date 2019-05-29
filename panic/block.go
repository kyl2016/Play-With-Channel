package main

import "log"

// fatal error: all goroutines are asleep - deadlock!
func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
		}
	}()

	ch := make(chan int)

	<-ch
}
