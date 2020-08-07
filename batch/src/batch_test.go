package src

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var count int32 = 0
var b = NewBatchImp(100, time.Second, 100, func(i interface{}) error {
	atomic.AddInt32(&count, int32(len(i.([]interface{}))))
	time.Sleep(time.Millisecond)
	return nil
})

func TestBatch_Add(t *testing.T) {
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

	fmt.Printf("concurrent=%d, amount per goroutine=%d, elaspsed: %fs, processed: %d\n", concurrent, amount, time.Since(start).Seconds(), count)
	time.Sleep(time.Second)
	if int(count) != concurrent*amount {
		t.Errorf("processed count is %d, should be %d", count, concurrent*amount)
	}
	// concurrent=4, amount per goroutine=1000000, elaspsed: 4.245187s
}

func TestBatch_Add_Small_Count(t *testing.T) {
	concurrent := 2
	amount := 101
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

	fmt.Printf("concurrent=%d, amount per goroutine=%d, elaspsed: %fs, processed: %d\n", concurrent, amount, time.Since(start).Seconds(), count)
}

func TestBatch_StopBeforeAdd(t *testing.T) {
	b = NewBatchImp(100, time.Second, 10, func(i interface{}) error {
		fmt.Println(i)
		time.Sleep(time.Millisecond)
		return nil
	})
	for i := 0; i < 10; i++ {
		b.Add(i)
		b.Stop()
	}
}

func TestBatch_ConcurrentStop(t *testing.T) {
	b.Add(1)

	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			b.Stop()
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestBatch_AddSlowly(t *testing.T) {
	b = NewBatchImp(100, time.Second, 1, func(i interface{}) error {
		fmt.Println(i)
		time.Sleep(time.Millisecond)
		return nil
	})
	for i := 0; i < 10; i++ {
		b.Add(i)
		time.Sleep(time.Second)
	}
	b.Stop()
}
