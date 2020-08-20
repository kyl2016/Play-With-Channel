package main

import (
	"github.com/Sirupsen/logrus"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(3 * time.Second)
		timeout <- true
	}()

	select {
	case i := <-ch:
		logrus.Info(i)
	case <-timeout:
		logrus.Info("timeout")
	}
}
