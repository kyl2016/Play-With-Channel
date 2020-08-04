package main

func main() {
	n := 1

	println(n << 1)
	println(n >> 31)
	println((n << 1) ^ (n >> 31))
}
