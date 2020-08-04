package main

func main() {
	r := test()
	println(r)
}

func test() (i int) {
	defer func() {
		i = i * 2
		println(i)
	}()

	println(i)

	return 3
}
