package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

var commands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

var Version = "0.1.0"

func main() {
	for i := 0; i < 100; i++ {
		Open("http://192.168.11.37:4000/#/")

		time.Sleep(5 * time.Second)
	}
}

func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}
