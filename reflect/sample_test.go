package performance

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestR(t *testing.T) {
	var err *os.LinkError
	tt := reflect.TypeOf(err).Elem()
	fmt.Printf("%+v\n", tt)

	tt = reflect.TypeOf(())
}
