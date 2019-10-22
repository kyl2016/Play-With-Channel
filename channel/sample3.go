package main

import (
	"fmt"
	"time"
)

func main() {
	numCh := make(chan struct{}, 1)
	abcCh := make(chan struct{}, 1)

	abcCh <- struct{}{}

	go func() {
		for i := 1; i <= 26; i++ {
			<-abcCh
			fmt.Println(i)
			numCh <- struct{}{}
		}
	}()

	words := "abcdefghijklmnopqrstuvwxyz"
	go func() {
		for _, w := range words {
			<-numCh
			fmt.Printf("%s\n", string(w))
			abcCh <- struct{}{}
		}
	}()

	time.Sleep(time.Second)
}
