package main

import (
	"fmt"
	"github.com/kyl2016/Play-With-Golang/basic/trace/stacktrace_from_runtime/errors"
	"strconv"
)

func atoi(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.Wrap(err)
	}
	return i, nil
}

func main() {
	_, err := atoi("f42")
	if err != nil {
		fmt.Println(err)
	}
}
