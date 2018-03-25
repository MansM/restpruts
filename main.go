package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	prometheus.MustRegister(newEvents)
	prometheus.MustRegister(requestedEvents)
	events = append(events, Event{ID: 1, Name: "Meetup"})
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
