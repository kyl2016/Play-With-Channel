package main

func main() {
	n := 4
	inc(&n)
	println(n)
}

func inc(x *int) {
	*x++
}

// show escape analysis:
// -m print optimization decisions
// $ go build -gcflags="-m" ./cpu.go
