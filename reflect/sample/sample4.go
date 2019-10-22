package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type User struct {
		Id   int
		Name string
		Age  int
	}

	var user User

	data := []byte(`
{
"id": 1,
"name": "kitty",
"age": 23
}
`)

	// panic: meet error:json: Unmarshal(non-pointer main.User)
	//err := json.Unmarshal(data, user)

	err := json.Unmarshal(data, &user)
	if err != nil {
		panic("meet error:" + err.Error())
	}

	fmt.Printf("%+v", user)
}
