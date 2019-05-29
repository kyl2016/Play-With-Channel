package main

import (
	"fmt"
	"time"
)

func counting(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}

	close(c)
}

func main() {
	//i := 0
	//for {
	//	i++
	//}

	msg := "Starting main"
	fmt.Println(msg)
	bus := make(chan int)
	msg = "starting a gofunc"

	ch2 := make(chan int)

	go func() {
		var i = 0
		for {
			i++
			time.Sleep(time.Second * 10)
			ch2 <- 1
		}
	}()

	go counting(bus)
	for count := range bus {
		fmt.Println("count:", count)

		<-ch2
	}

}
