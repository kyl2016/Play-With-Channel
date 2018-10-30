package main

import (
	"github.com/Sirupsen/logrus"
	"time"
)

func main() {
	ch := make(chan int, 10)
	ch2 := make(chan int, 10)

	for {
		select {
		case ch <- 1:
			logrus.Info("Write 1")
			time.Sleep(0.2e9)
		case ch2 <- 2:
			logrus.Info("Write 2")
			time.Sleep(0.2e9)
		default:
			logrus.Info("Channel is full.")
			return
		}
	}
}
