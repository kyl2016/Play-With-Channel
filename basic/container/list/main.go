package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(&list.Element{Value: "1"})
	l.PushBack(&list.Element{Value: "2"})

	for {
		e := l.Front()
	if e != nil {
		fmt.Println(e.Value)
		l.Remove(e)
	} else {
		break
	}
	}

	fmt.Println(l)
}
