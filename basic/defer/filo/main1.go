package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("first defer")

	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in  for [%d]\n", i) // i作为函数参数传入，会预计算
	}

	defer fmt.Println("last defer")
}
