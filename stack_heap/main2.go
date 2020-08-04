package main

func main() {
	n := answer()
	println(*n / 2)
}

func answer() *int {
	x := 42
	return &x
}

// escapes to the heap:
// $ go build -gcflags="-m -l" ./main2.go
