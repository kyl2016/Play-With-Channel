package filo

import "fmt"

func main() {
	defer fmt.Println("first defer")

	for i := 0; i < 3; i++ {
		defer func(i2 int) {
			fmt.Printf("defer in  for [%d]\n", i2)
		}(i)
	}

	defer fmt.Println("last defer")
}
