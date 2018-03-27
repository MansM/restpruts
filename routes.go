package main

import (
	"net/http"
)

//Route contains the route to a handler
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes list of route structs
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetEvents",
		"GET",
		"/event",
		GetEvents,
	},
	Route{
		"GetEvent",
		"GET",
		"/event/{id}",
		GetEvent,
	},
	Route{
		"CreateEvent",
		"POST",
		"/event",
		CreateEvent,
	},
}
