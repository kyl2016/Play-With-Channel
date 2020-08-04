package test

import (
	"flag"
	"fmt"
	"log"
	"testing"
)

var para = flag.String("para", "default", "test para")

func TestT(t *testing.T) {
	fmt.Println("para: ",*para)
	if *para != "testPara"{
		t.Error("error")
	}
	fmt.Println("test")
}

func BenchmarkT(t *testing.B) {
	t.Error("error")
	fmt.Println("test benchmark ", t.N)
}

func ExampleT(){
	fmt.Println("example")

	log.Fatal("fatal error")
}