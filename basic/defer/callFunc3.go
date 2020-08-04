package main

import "fmt"

func main() {
	var d *Data3
	d = new(Data3)
	d.Max = 1
	fmt.Printf("%p %v\n", d, d)

	defer deferExec3(d.Println)

	d = new(Data3)
	d.Max = 2
	fmt.Printf("%p %v\n", d, d)
}

type Data3 struct {
	Max int
}

func (t *Data3) Println() {
	fmt.Printf("%p %v\n", t, t)
}

func deferExec3(f func()) {
	f()
}
