//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//// 一个栅栏接口。
//type Barrier interface {
//	Wait()
//}
//
//// 创建栅栏对象。
//func NewBarrier(n int) Barrier {
//	// 这里需要填入代码。
//	b := barrier{max: n}
//	b.chans = make([]chan interface{}, n-1)
//	b.locker = &sync.Mutex{}
//	return &b
//}
//
//// 栅栏的实现类型。
//type barrier struct {
//	// 这里需要填入代码。
//	chans  []chan interface{}
//	count  int
//	max    int
//	locker sync.Locker
//}
//
//func (b *barrier) Wait() {
//	// 这里需要填入代码。
//
//	//b.locker.Lock()
//	if b.count < b.max-1 {
//		b.count++
//		ch := make(chan interface{})
//		b.chans[b.count-1] = ch
//		//b.locker.Unlock()
//
//		<-ch
//	} else {
//		for _, ch := range b.chans {
//			close(ch)
//		}
//		//b.locker.Unlock()
//	}
//}
//
//// 测试代码。
//func main() {
//	num := 10
//	// 创建栅栏值。
//	b := NewBarrier(num)
//	// 需要达到的效果：
//	//   前 9 个 goroutine 调用 Wait 方法时被阻塞；
//	//   第 10 个 goroutine 调用 Wait 方法后，所有 goroutine 全部被唤醒。
//	for i := 0; i < num; i++ {
//		go func(i int) {
//			fmt.Printf("Wait[%d]\n", i)
//			b.Wait()
//			fmt.Printf("Done[%d]\n", i)
//		}(i)
//	}
//	time.Sleep(time.Second * 2)
//}
