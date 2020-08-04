package main

import (
	"net/url"
	"github.com/gorilla/websocket"
	"fmt"
)

func main() {
	u := url.URL{
		Scheme:   "ws",
		Host:     "localhost:4000",
		Path:     "ping",
		RawQuery: "",
	}

	c, _, err := websocket.DefaultDailer.Dail(u.String(), nil)
	if err != nil {
		panic(err)
	}

	go func(){
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				break
			}

			fmt.Println(string(message))
		}
	}

	c.WriteMessage("send message from client")

	select{}
}
