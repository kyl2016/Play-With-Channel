package main

import "github.com/Sirupsen/logrus"

func main() {
	ch := make(chan int, 1)
	ch <- 2

	select {
	case ch <- 1:
		logrus.Info("Write ok.")
	default:
		logrus.Info("Channel is full.")
	}

	for {
		select {
		case i := <-ch:
			logrus.Info("Read ", i)
		default:
			logrus.Info("Channel is empty.")
			return
		}
	}
}
