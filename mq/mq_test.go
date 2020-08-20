package mq

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestMQ_Unsub(t *testing.T) {
	mq := NewMQ()
	id, ch := mq.Sub("t1")
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	mq.Pub("t1", 1)
	mq.Pub("t1", 2)
	mq.Unsub(id, "t1")
	mq.Pub("t1", 3)
	time.Sleep(time.Second)
}

func TestMQ_Unsub_Concurrent(t *testing.T) {
	var count int32
	mq := NewMQ()
	id, ch := mq.Sub("t1")
	go func() {
		for range ch {
			atomic.AddInt32(&count, 1)
			//fmt.Println(i)
		}
		fmt.Println(count)
	}()

	go func() {
		time.Sleep(time.Second * 2)
		mq.Unsub(id, "t1")
	}()

	for j := 0; j < 100; j++ {
		go func() {
			for i := 0; i < 10000; i++ {
				mq.Pub("t1", i)
			}
		}()
	}

	time.Sleep(time.Second * 15)
}

func TestMQ_Pub(t *testing.T) {
	var count int32
	mq := NewMQ()
	_, ch := mq.Sub("t1")
	go func() {
		for range ch {
			atomic.AddInt32(&count, 1)
		}
	}()

	for i := 0; i < 10; i++ {
		start := time.Now()
		wg := sync.WaitGroup{}
		for j := 0; j < 3000; j++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 1000; i++ {
					mq.Pub("t1", i)
				}
				wg.Done()
			}()
		}
		wg.Wait()
		fmt.Println(time.Since(start), count)
	}
}
