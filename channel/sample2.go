package main

import (
	"fmt"
	"time"
)

// 一个栅栏接口。
type Barrier interface {
	Wait()
}

// 创建栅栏对象。
func NewBarrier(n int) Barrier {
	// 这里需要填入代码。
	if n < 1 {
		n = 1
	}
	b := &barrier{
		max:    n,
		waitCh: make(chan struct{}),
		passCh: make(chan struct{}),
	}
	go b.watch()
	return b
}

// 栅栏的实现类型。
type barrier struct {
	// 这里需要填入代码。
	max    int
	waitCh chan struct{}
	passCh chan struct{}
}

func (b *barrier) watch() {
	for i := 0; i < b.max; i++ {
		<-b.waitCh
	}
	for j := 0; j < b.max; j++ {
		b.passCh <- struct{}{}
	}
}

func (b *barrier) Wait() {
	// 这里需要填入代码。
	b.waitCh <- struct{}{}
	<-b.passCh
}

// 测试代码。
func main() {
	num := 10
	// 创建栅栏值。
	b := NewBarrier(num)
	// 需要达到的效果：
	//   前 9 个 goroutine 调用 Wait 方法时被阻塞；
	//   第 10 个 goroutine 调用 Wait 方法后，所有 goroutine 全部被唤醒。
	for i := 0; i < num; i++ {
		go func(i int) {
			fmt.Printf("Wait[%d]\n", i)
			b.Wait()
			fmt.Printf("Done[%d]\n", i)
		}(i)
	}
	time.Sleep(time.Second * 2)
}
