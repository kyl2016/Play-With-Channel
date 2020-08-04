package routers

import "github.com/gorilla/mux"

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle(
		"/test/hello",
		negroni.New(
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}
