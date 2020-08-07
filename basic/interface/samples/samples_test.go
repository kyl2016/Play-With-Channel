package samples

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var i interface{}
	i = []interface{}{}
	fmt.Println(i)
}
