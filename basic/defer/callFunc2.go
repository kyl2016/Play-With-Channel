package main

func main() {
	var d *Data
	defer deferExec2(d.S.F)
	d = new(Data)
}

type Sub struct {
	F func()
}

type Data struct {
	Max int
	S   Sub
}

func (t *Data) Println() {
	println(t.Max)
}

func deferExec2(f func()) {
	f()
}
