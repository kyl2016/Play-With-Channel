/* EchoChanServer
 */
package main

import (
	"fmt"
	"old/netchan"
	"os"
	"strconv"
)

var count int = 0

func main() {
	exporter := netchan.NewExporter()
	err := exporter.ListenAndServe("tcp", ":2345")
	checkError(err)

	echo := make(chan string)
	exporter.Export("echo", echo, netchan.Send)
	for {
		sCount := strconv.Itoa(count)
		lock := make(chan string)
		go handleSession(exporter, sCount, lock)

		<-lock
		echo <- sCount
		count++
		exporter.Drain(-1)
	}
}

func handleSession(exporter *netchan.Exporter, sCount string, lock chan string) {
	echoIn := make(chan string)
	exporter.Export("echoIn"+sCount, echoIn, netchan.Send)

	echoOut := make(chan string)
	exporter.Export("echoOut"+sCount, echoOut, netchan.Recv)
	fmt.Println("made " + "echoOut" + sCount)

	lock <- "done"

	for {
		s := <-echoOut
		echoIn <- s
	}
	// should unexport net channels
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
