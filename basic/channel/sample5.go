package main

func main() {
	for i := 0; i < 5; i++ {
		go println(i)
	}
}
