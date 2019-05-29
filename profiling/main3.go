package main

import (
	"context"
	"fmt"
	. "github.com/pkg/profile"
	"math/rand"
	"time"
)

func main() {
	defer Start().Stop()

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	test(ctx)
}

func test(c context.Context) {
	i := 0
	j := 0

	ch := make(chan int)

	go func() {
		m := map[int]int{}
		for {
			i++
			m[i] = i

			select {
			case _ = <-ch:
				//fmt.Println(r)
			}
		}
	}()

	go func() {
		m := map[int]int{}
		for {
			select {
			case ch <- j:
			case <-time.After(time.Duration(rand.Intn(500)) * time.Millisecond):
			}

			j++
			m[i] = i
		}
	}()
	select {
	case <-c.Done():
		fmt.Println("done, i", i, "j", j)
		return
	}
}
