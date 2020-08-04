package main

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/cloudfoundry/gosigar"
	"fmt"
)

func main() {
	println(primitive.NewObjectID().Hex())
	fmt.Println("hello")

	_ = sigar.Mem{}
}
