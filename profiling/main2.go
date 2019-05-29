package main

import (
	"context"
	"fmt"
	. "github.com/pkg/profile"
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
	go func() {
		m := map[int]int{}
		for {
			i++
			m[i] = i
		}
	}()
	go func() {
		m := map[int]int{}
		for {
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
