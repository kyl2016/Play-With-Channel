package main

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/kyl2016/Play-With-Golang/trace/go-errors/crashy"
)

func main() {
	err := crashy.Crash()
	if err != nil {
		fmt.Println(err.(*errors.Error).ErrorStack())
	}
}
