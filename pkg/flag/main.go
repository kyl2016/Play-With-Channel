package main

import (
	"flag"
	"fmt"
)

var name string
var age = flag.Int("age", 0, "input age")

func main() {
	flag.StringVar(&name, "name", "default", "hello")
	flag.Parse()

	PrintInfo(fmt.Sprintln("hello", name, *age))
}

// 执行
// go run ./cpu.go ./output.go -name Bob -age 10
// go run . -name Bob -age 10

// 查看输入参数
// go run ./cpu.go ./output.go --help

// go build ./cpu.go ./output.go
// ./main --help

// install, 会在 ~/go/bin 下面生成 flag
// go install .
