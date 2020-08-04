package main

import (
	"github.com/codegangsta/negroni"
	"github.com/kyl2016/Play-With-Golang/pkg/jwt/routers"
	"net/http"
)

func main() {
	settings.Init()
	router := routers.InitRouters()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
