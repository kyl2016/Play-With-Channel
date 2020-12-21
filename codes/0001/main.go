package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			<-time.After(time.Second)
			func() {
				defer func() {
					if p := recover(); p != nil {
						fmt.Println(p)
					}
				}()
				proc()
			}()
		}
	}()

	select {}
}

func proc() {
	panic("ok")
}
