package samples

import (
	"fmt"
	"testing"
	"time"
)

// after close(ch), set ch=nil to forbidden receive false
func ExampleClose(t *testing.T) {
	ch := make(chan int, 1)
	ch2 := make(chan int)
	go func() {
		for {
			select {
			case _, ok := <-ch:
				fmt.Println("1", ok)
				ch = nil
			case _, ok := <-ch2:
				fmt.Println("2", ok)
			}
		}
	}()

	ch <- 1
	close(ch)
	time.Sleep(time.Millisecond)

	// output:
	// 1 true
}
