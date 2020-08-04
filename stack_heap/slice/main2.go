package main

func main() {
	s := make([]byte, 10)
	write(s)
	println(len(s))
}

func write(s []byte) {
	s[0] = 'a'
}
