package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Second * 2)
		cancel()
		println("canceled")
	}()

	for {
		if canceled(ctx) {
			println("stop doing")
			break
		}

		time.Sleep(time.Second)
	}

	fmt.Println("error:", ctx.Err())

	println("finished")
}

func canceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
