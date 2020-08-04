package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Minute)
	}()

	log.Println(http.ListenAndServe("localhost:8082", nil))
}
