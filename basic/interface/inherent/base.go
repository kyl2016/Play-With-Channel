package inherent

import "fmt"

type Base struct {
	Name string
}

func (b *Base) Step1() {
	fmt.Println("base step 1")
}

func (b *Base) Step2() {
	fmt.Println("base step 2")
}

func (b *Base) Step3() {
	fmt.Println("base step 3")
}

func (b *Base) Test() {
	b.Step1()
	b.Step2()
	b.Step3()
}
