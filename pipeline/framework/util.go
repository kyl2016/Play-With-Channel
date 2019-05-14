package framework

import "fmt"

func writeLog(ch chan string, content string) {
	select {
	case ch <- content:
	default:
		fmt.Printf("LogCh is full, can't write content:%s\n", content)
	}
}
