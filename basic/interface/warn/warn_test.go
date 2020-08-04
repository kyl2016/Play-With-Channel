package warn

import (
	"fmt"
	"testing"
)

type Pet interface {
	SetName(name string)
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func TestIsNil(t *testing.T) {
	var dog Dog
	var pet Pet = &dog
	if pet == nil {
		println("pet is nil")
	}
}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("nil interface")
		return
	}
	fmt.Printf("non-nil interface %v\n", x)
}

func TestFoo(t *testing.T) {
	var x *int = nil
	Foo(x)
	Foo(nil)
}
