package main

func main() {
	i := 100
	for {
		i += i / 4
		println(i)

		if i <= 0 {
			println("overflow")
			break
		}
	}
}
