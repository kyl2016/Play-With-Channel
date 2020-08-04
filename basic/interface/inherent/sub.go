package inherent

import "fmt"

type Sub struct {
	Base
}

func (s *Sub) Step2() {
	fmt.Println("Sub step 2")
}

func (b *Sub) Test() {
	b.Base.Test()
}
