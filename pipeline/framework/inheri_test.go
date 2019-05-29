package framework

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {


	var s interface{}
	s = Status{}
	T(s.(ErrorInfo))
}

func T(info ErrorInfo) {
	fmt.Printf("%v", info)
}