package main

import "fmt"

func main() {
	c := Gen(2, 3)
	out := Sq(c)

	fmt.Println(<-out)
	fmt.Println(<-out)

	for n := range Sq(Sq(Gen(2, 3))) {
		fmt.Println(n)
	}
}

func Gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func Sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}
