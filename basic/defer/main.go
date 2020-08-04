package main

import "fmt"

type number int

func (n number) print()         { fmt.Println(n) }
func (n *number) printPointer() { fmt.Println(*n) }

func main() {
	var n number
	defer n.print()
	defer n.printPointer()
	defer func() { n.print() }()
	defer func() { n.printPointer() }()

	n = 3
}
