package main

import (
	"log"
	"net/http"
)

func main() {
	events = append(events, Event{ID: 1, Name: "Meetup"})
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
