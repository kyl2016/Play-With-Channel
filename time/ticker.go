package main

import "time"

func main() {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			println(time.Now().String())
		}
	}
}
