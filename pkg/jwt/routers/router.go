package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRouters(router)
	router = SetAuthenticationRouters(router)
	return router
}
