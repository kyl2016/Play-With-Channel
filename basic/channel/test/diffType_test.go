package test

import "testing"

func TestDiffType(t *testing.T) {
	ch := make(chan interface{})
	ch <- struct{ name string }{name: "kitty"}

	printCh(ch)
}

func printCh(ch chan interface{}) {
	data := <-ch
	println(data)
}
