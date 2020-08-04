package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, _ := os.Create("lookup")

	p := pprof.Lookup("goroutine")
	defer p.WriteTo(f, 0)

	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(time.Microsecond)
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Millisecond)
}
