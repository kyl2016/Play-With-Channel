package filo

import "fmt"

func main() {
	defer fmt.Println("first defer")

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Printf("defer in  for [%d]\n", i)
		}()
	}

	defer fmt.Println("last defer")
}
