package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter blah
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	//fmt.Printf(string(router.GetRoute(routes[0].Name)))
	return router
}
