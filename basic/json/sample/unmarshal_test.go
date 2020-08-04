package sample

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	j := `{"key":"value"}`
	//var m map[string]interface{}
	//m = map[string]interface{}{}
	var m interface{}
	m = nil

	err := json.Unmarshal([]byte(j), &m)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
