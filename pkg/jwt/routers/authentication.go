package routers

import "github.com/gorilla/mux"

func SetAuthenticationRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/token-auth", controllers.Login).Methods("POST")
	router.HandleFunc("/refresh-token-auth", negroni.HandlerFunc(controllers.RefreshToken)).Methods("GET")
	router.HandleFunc("/logout", negroni.New(negroni.HandlerFunc(authentication.RequireTokenAuthentication), negroni.HandlerFunc(controllers.Logout))).Methods("GET")
	return router
}
