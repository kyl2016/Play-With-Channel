package main

import "sync/atomic"

func main() {
	var n int32 = 0
	m := atomic.AddInt32(&n, 1)
	println(n, m)
	m = atomic.AddInt32(&n, 10)
	println(n, m)
}
