package main

import "fmt"

type T struct {
	num int
}

func (t *T) Increase() {
	t.num++
}

func (t T) F1() {
	fmt.Println(t.num)
}

func (t *T) F2() {
	fmt.Println(t.num)
}

func main() {
	var t T
	t.F1()
	t.F2()

	t.Increase()

	(&t).F1()
	(&t).F2()
}
