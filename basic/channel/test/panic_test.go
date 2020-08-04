package test

import "testing"

func Test_SendOnClosedChannel(t *testing.T) {
	ch := make(chan int, 1)
	close(ch)
	ch <- 1
}

func Test_CloseOfClosedChannel(t *testing.T) {
	ch := make(chan int)
	close(ch)
	close(ch)
}
