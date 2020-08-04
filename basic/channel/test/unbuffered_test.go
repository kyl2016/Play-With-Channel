package test

import (
	"testing"
	"time"
)

func TestModifyWhenBlocked(t *testing.T) {
	ch := make(chan int)

	i := 1

	go func() {
		ch <- i
	}()

	time.Sleep(time.Millisecond * 10)

	println(<-ch)
}

func TestModifyWhenBlocked2(t *testing.T) {
	ch := make(chan int)

	i := 1

	go func() {
		ch <- i
	}()

	time.Sleep(time.Millisecond * 10)

	i = 2

	println(<-ch)
	// 虽然是 unbuffered，但 i 的值已经 copy 到 ch 的 sendq 中 g 的 elem 中了
}

func TestModifyWhenBlocked3(t *testing.T) {
	ch := make(chan *int)

	i := 1

	go func() {
		ch <- &i
	}()

	time.Sleep(time.Millisecond * 10)
	i = 3

	println(*<-ch)
}
