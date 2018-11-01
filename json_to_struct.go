package main

import (
	"github.com/Sirupsen/logrus"
	"gosrc/src/encoding/json"
)

type Child struct {
	Name string `json:"name"`
}

type Person struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Desc     string  `json:"desc"`
	Children []Child `json:children`
}

func main() {
	jsonStr := `
{
	"name": "kitty",
	"age": 22,
	"desc": "She is a teacher.",
	"children": [
		{"name": "bob"}
	]
}
`

	var person Person

	json.Unmarshal([]byte(jsonStr), &person)

	logrus.Info(person)
}
