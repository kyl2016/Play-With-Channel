package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	ch := make(chan string)

	go server(ch)

	ch <- "work1"
	ch <- "work2"

	//go func() {
	//	panic("other goroutine")
	//}()

	time.Sleep(500 * time.Millisecond)

	fmt.Println("finished")
}

func server(workChan <-chan string) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic:", err)
		}
	}()

	do(work)
}

func do(work string) {

	if time.Now().Nanosecond()%2 == 0 {
		panic(fmt.Sprintf("%s %s", work, "failed"))
	}

	log.Println(work, "done.")
}
