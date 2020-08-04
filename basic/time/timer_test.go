package main

import (
	"fmt"
	"time"
)

func ExampleTimer_Simple() {
	t := time.NewTimer(time.Millisecond * 100)
	<-t.C
	t.Stop()
	t.Reset(time.Millisecond * 100)
	<-t.C

	// Output:
}

func ExampleTimer() {
	t := time.NewTimer(time.Second)
	go func() {
		time.Sleep(time.Second * 3)
		ok := false
		for !ok {
			ok = t.Stop()
			if !ok {
				fmt.Println("Stop:", ok)
			}
		}
	}()

	//fmt.Println(<-t.C)

	for {
		ti := <-t.C
		fmt.Println(ti)
		t.Reset(time.Second)
	}

	//for ti := range t.C {
	//	fmt.Println(ti)
	//}

	// Output:
}
