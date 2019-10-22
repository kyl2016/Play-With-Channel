package main

import "fmt"

type Animal interface {
	SetName(name string)
	ShowInfo()
}

type Cat struct {
	name string
}

func (c Cat) SetName(name string) {
	c.name = name
}

func (c Cat) ShowInfo() {
	println(c.name)
}

type Dog struct {
	name string
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func (d *Dog) ShowInfo() {
	println(d.name)
}

func main() {
	var a Animal
	a = Cat{}
	a.SetName("")
}

type People struct{}

func (p People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teachershowB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
