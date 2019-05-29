package main

func main() {
	bytes := []byte("5c3ed3444844b428a65dc0b6")
	println(len(bytes))

	bytes = []byte("gc3ed3444844b428a65dc0b6")
	println(len(bytes))
}
