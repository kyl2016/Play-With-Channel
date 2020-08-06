package main

import "net"

func main() {
	connCh := make(chan net.Conn, 10)
	outCh := make(chan []byte, 100)
	for conn := range connCh {
		go process(conn, outCh)
	}
}

func process(conn net.Conn, outCh chan []byte) {
	for {
		var buffer []byte
		_, err := conn.Read(buffer)
		if err != nil {
			return
		}
		outCh <- buffer
	}
}

func read() []byte {
	return []byte("test")
}
