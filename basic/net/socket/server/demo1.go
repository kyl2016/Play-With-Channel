package main

import(
	"net"
	"bufio"
	"strings"
)

func main() {
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		println(message)
		upper := strings.ToUpper(message)
		conn.Write([]byte(upper+"\n"))
	}
}