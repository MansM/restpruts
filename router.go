package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewRouter blah
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.Methods(http.MethodGet).Path("/metrics").Handler(promhttp.Handler())
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
