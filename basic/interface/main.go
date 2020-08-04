package main

import (
	"fmt"
	"github.com/kyl2016/Play-With-Golang/basic/interface/inherent"
)

func main() {
	sub := inherent.Sub{}
	sub.Name = "sub"

	sub.Test()

	fmt.Println("")

	sub.Base.Test()
}
