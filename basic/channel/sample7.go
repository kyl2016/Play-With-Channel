package main

func main() {
	ch := make(chan int, 1)
	close(ch)

	select {
	case ch <- 1:
	default:
		println("default")
	}
}
