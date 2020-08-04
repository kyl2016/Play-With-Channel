package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	for i := 0; i < 100000; i++ {
		go func() {
			for {
				fmt.Println("hello world")
			}
		}()
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
