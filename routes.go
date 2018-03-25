package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

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
