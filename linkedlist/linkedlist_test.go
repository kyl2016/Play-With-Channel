package linkedlist

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	node1 := &LinkedNode{Data: 1, Next: nil}

	list := LinkedList{node1}
	Reverse(&list)
	print(list)

	list.Node = nil
	Reverse(&list)
	print(list)
}

func TestReverse2(t *testing.T) {
	node1 := &LinkedNode{Data: 1, Next: nil}

	node2 := &LinkedNode{Data: 2, Next: nil}
	node1.Next = node2

	list := LinkedList{node1}
	list.Node = node1

	print(list)
	Reverse(&list)
	print(list)
}

func TestReverse3(t *testing.T) {
	node1 := &LinkedNode{Data: 1, Next: nil}
	list := LinkedList{node1}

	node2 := &LinkedNode{Data: 2, Next: nil}
	node1.Next = node2
	list.Node = node1

	node3 := &LinkedNode{Data: 3, Next: nil}
	node2.Next = node3
	print(list)
	Reverse(&list)
	print(list)
}

func print(list LinkedList) {
	fmt.Println()

	if list.Node == nil {
		fmt.Println("nil")
		return
	}

	for {
		if list.Node == nil {
			break
		}

		fmt.Printf("%v,", list.Node.Data)
		list.Node = list.Node.Next
	}

	fmt.Println("")
}
