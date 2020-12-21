package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":4001", &MyHandler{})
}

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(*r.URL)
	w.Write([]byte("hello"))
}
