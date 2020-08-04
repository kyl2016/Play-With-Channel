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
		for i := 1; i <= 26; i += 2 {
			<-abcCh
			fmt.Printf("%d%d", i, i+1)
			numCh <- struct{}{}
		}
	}()

	words := "abcdefghijklmnopqrstuvwxyz"
	go func() {
		for i := 0; i < 26; i += 2 {
			<-numCh
			fmt.Printf("%s%s", string(words[i]), string(words[i+1]))
			abcCh <- struct{}{}
		}
	}()

	time.Sleep(time.Second)
}
