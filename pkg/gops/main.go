package main

import (
	"github.com/google/gops/agent"
	"log"
	"time"
)

func main() {
	startGoroutine()
	startGoroutine()
	startGoroutine()

	if err := agent.Listen(agent.Options{
		Addr:            "",
		ConfigDir:       "",
		ShutdownCleanup: false,
	}); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour)
}

func startGoroutine() {
	go func() {
		time.Sleep(time.Minute)
	}()
}
