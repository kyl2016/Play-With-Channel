package samples

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

//

func TestSelectCase(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		fmt.Println("send to ch1")
		ch1 <- 1
		fmt.Println("close ch1")
		close(ch1)
	}()

	go func() {
		fmt.Println("send to ch2")
		ch2 <- 2
		time.Sleep(time.Second * 1)
		fmt.Println("close ch2")
		close(ch2)
	}()

	r := merge(ch1, ch2)
	for i := range r {
		fmt.Println(i)
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := merge(asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
		for range c {

		}
	}
}

func BenchmarkMergeReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := mergeReflect(asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
		for range c {

		}
	}
}

func BenchmarkMergeRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := mergeRec(asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
		for range c {

		}
	}
}

// Unfortunately, that simply gives a single and pretty meaningless measure, mixing all performances.
func BenchmarkMergeAll(b *testing.B) {
	merges := []func(...<-chan int) <-chan int{
		merge,
		mergeReflect,
		mergeRec,
	}
	for _, merge := range merges {
		for i := 0; i < b.N; i++ {
			c := merge(asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
			for range c {

			}
		}
	}
}

// create subbenchmarks by calling testing.B.Run
func BenchmarkMergeAll2(b *testing.B) {
	merges := []struct {
		name string
		fun  func(...<-chan int) <-chan int
	}{
		{"goroutines", merge},
		{"reflect", mergeReflect},
		{"recursion", mergeRec},
	}
	for _, merge := range merges {
		b.Run(merge.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c := merge.fun(asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
				for range c {
				}
			}
		})
	}
}

func BenchmarkMergeAll_MultiChans(b *testing.B) {
	merges := []struct {
		name string
		fun  func(...<-chan int) <-chan int
	}{
		{"goroutines", merge},
		{"reflect", mergeReflect},
		{"recursion", mergeRec},
	}

	var chans []<-chan int
	for i := 0; i < 10240; i++ {
		chans = append(chans, asChan(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	}

	for _, merge := range merges {
		b.Run(merge.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c := merge.fun(chans...)
				for range c {
				}
			}
		})
	}
}

//                                                         循环执行次数   每次循环花费
//BenchmarkMergeAll_MultiChans/goroutines
//BenchmarkMergeAll_MultiChans/goroutines-4         	     284	   4263334 ns/op
//BenchmarkMergeAll_MultiChans/reflect
//BenchmarkMergeAll_MultiChans/reflect-4            	       1	49502125370 ns/op
//BenchmarkMergeAll_MultiChans/recursion
//BenchmarkMergeAll_MultiChans/recursion-4          	     120	  17196815 ns/op

func merge(chans ...<-chan int) <-chan int {
	resultCh := make(chan int, 100)
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch <-chan int) {
			defer wg.Done()
			for item := range ch {
				resultCh <- item
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()
	return resultCh
}

// This is a nice solution that avoids using reflection and also reduces the number
// of goroutines needed. On the other hand, it uses more channels than before.
func mergeRec(chans ...<-chan int) <-chan int {
	switch len(chans) {
	case 0:
		c := make(chan int)
		close(c)
		return c
	case 1:
		return chans[0]
	default:
		m := len(chans) / 2
		return mergeTwo(
			mergeRec(chans[:m]...),
			mergeRec(chans[m:]...))
	}
}

func mergeTwo(ch1 <-chan int, ch2 <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for ch1 != nil && ch2 != nil {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				c <- v
			case v, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}

func mergeReflect(chans ...<-chan int) <-chan int {
	resultCh := make(chan int, 10)

	var cases []reflect.SelectCase
	for _, ch := range chans {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	go func() {
		defer close(resultCh)

		for {
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				if len(cases) == 0 {
					return
				}
				continue
			}
			resultCh <- int(v.Int())
		}
	}()

	return resultCh
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for _, v := range vs {
			c <- v
			time.Sleep(time.Millisecond)
		}
	}()
	return c
}

// go test -bench=. -run=BenchmarkMerge
