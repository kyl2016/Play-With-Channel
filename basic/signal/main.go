package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	fmt.Println(os.Getpid())

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		signal := <-ch
		fmt.Println("signal:", signal)
		wg.Done()
	}()

	wg.Wait()
}
