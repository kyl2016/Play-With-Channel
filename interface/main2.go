package main

type Fragment interface {
	Exec()
	Print()
}

type GetPodAction struct {
	Num int
}

func (g GetPodAction) Exec() {
	g.Num++
}

func (g GetPodAction) Print() {
	println(g.Num)
}

func main() {
	var f1 Fragment = GetPodAction{}
	println(f1)
	f1.Exec()
	f1.Print()
	var f2 Fragment = &GetPodAction{}
	println(f2)
	f2.Exec()
	f2.Print()
	var f3 Fragment = new(GetPodAction)
	println(f3)
	f3.Exec()
	f3.Print()
}
