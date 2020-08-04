package main

import (
	"time"
)

func ExampleTicker() {
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(time.Second * 5)
		ticker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			println(time.Now().String())
		}
	}

	// Output:
}
