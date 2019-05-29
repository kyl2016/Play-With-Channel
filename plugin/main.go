package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("main.so")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("Hello")
	if err != nil {
		panic(err)
	}

	fmt.Println(p)

	f.(func(s string))("test")
}

// go build -buildmode=plugin ./pkg/main.go
// go run main.go