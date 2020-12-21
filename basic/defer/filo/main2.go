package main

import "fmt"

func main() {
	defer fmt.Println("first defer")

	for i := 0; i < 3; i++ {
		defer func() {
			// 闭包，i 引用了外部变量，拷贝的是函数指针，不会立即求值
			fmt.Printf("defer in  for [%d]\n", i)
		}()
	}

	defer fmt.Println("last defer")
}
