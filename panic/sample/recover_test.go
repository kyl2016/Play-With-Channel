package sample

import (
	"fmt"
	"testing"
	"time"
)

func TestRecover(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("inner goroutine:", r)
			}
		}()

		panic("test")
	}()

	time.Sleep(time.Millisecond)
}
