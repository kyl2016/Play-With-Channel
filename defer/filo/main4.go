package filo

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("first defer")

	for i := 0; i < 3; i++ {
		go func() {
			defer fmt.Printf("defer in  for [%d]\n", i)
		}()
	}

	time.Sleep(time.Millisecond * 100)

	defer fmt.Println("last defer")
}
