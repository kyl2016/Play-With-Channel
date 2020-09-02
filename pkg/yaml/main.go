package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}
	var c Config
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}

type Config struct {
	DB struct {
		Url   string
		Name  string
		Table string
	}
}
