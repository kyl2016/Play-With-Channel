package test

import (
	"testing"
)

func Test_SendToNil(t *testing.T) {
	var ch chan int
	ch <- 1
}

func Test_ReceiveFromNilChannel(t *testing.T) {
	var ch chan int
	<-ch
}
