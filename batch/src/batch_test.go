package src

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var count = 0

func TestBatch_Add(t *testing.T) {
	b := NewBatch(100, time.Second, func(i []interface{}) error {
		count += len(i)
		return nil
	})

	concurrent := runtime.NumCPU()
	amount := 1000000
	wg := sync.WaitGroup{}

	start := time.Now()
	for i := 0; i < concurrent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < amount; j++ {
				b.Add(j)
			}
		}()
	}

	wg.Wait()
	b.Stop()

	fmt.Printf("concurrent=%d, amount per goroutine=%d, elaspsed: %fs\n", concurrent, amount, time.Since(start).Seconds())
	// concurrent=4, amount per goroutine=1000000, elaspsed: 4.245187s
}
