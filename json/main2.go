package main

import (
	"encoding/json"
	"fmt"
)
type RawMessage []byte

type Result struct {
	Attr *RawMessage
}

func main() {
	attr := RawMessage("{}")
	attr = nil
	r := Result{Attr:&attr}
	r = Result{}

	buffer, _ := json.Marshal(&r)
	fmt.Println(string(buffer))
}
