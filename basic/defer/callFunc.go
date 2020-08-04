package main

func main() {
	var d *People
	defer deferExec(d.Println)
	d = new(People)
}

type People struct {
	Age uint32
}

func (t *People) Println() {
	println(t.Age)
}

func deferExec(f func()) {
	f()
}
