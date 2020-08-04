package main

func f(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func main() {
	c1 := f(0)
	c2 := f(0)
	println(c1())
	println(c2())
}
