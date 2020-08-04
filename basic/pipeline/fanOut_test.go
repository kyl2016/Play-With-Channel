package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestFanOut(t *testing.T) {
	in := Gen(2, 3)

	c1 := Sq(in)
	c2 := Sq(in)

	for n := range merge1(c1, c2) {
		fmt.Println(n)
	}
}

func merge1(cs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(len(cs))
	for _, c := range cs {
		go func() {
			for r := range c {
				out <- r
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func merge2(cs ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	output := func(c <-chan int) {
		for r := range c {
			out <- r
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
