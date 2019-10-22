package main

import "fmt"

type People struct{}

func (p People) ShowA() {
	fmt.Println("showA")
	(&p).ShowB()
	p.ShowB()
}

func (p People) ShowB() {
	fmt.Println("showB")
}

func (p *People) ShowC() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
	//t.ShowB()
}
