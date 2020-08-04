package main

import "fmt"

func main() {
	arr := []int{1,2,3}
	for i := range arr {
		defer func() {
			fmt.Println("defer") // effect ONLY after return
		}()

		fmt.Println(i)
	}
}