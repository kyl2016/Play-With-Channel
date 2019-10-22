package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("begin")
	v := string(Runcmd("git describe --abbrev=0 --tags", true))
	fmt.Println("tag:", v)
	fmt.Println("end")
}

func Runcmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			fmt.Println("error:", err)
			log.Fatal(err)
			//panic("some error found")
		}
		return out
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	return out
}
