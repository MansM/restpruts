package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Index bla
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

// GetEvents List all events
func GetEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(events)
}

// GetEvent Display just one event
func GetEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range events {

		if i, err := strconv.Atoi(params["id"]); err == nil {

			if i == item.ID {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	}
	requestedEvents.Inc()
	json.NewEncoder(w).Encode(&Event{})
}

// CreateEvent blah
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var eventn Event
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &eventn); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	events = append(events, Event{ID: eventn.ID, Name: eventn.Name})
	newEvents.Inc()
}

// DeleteEvent blah
func DeleteEvent(w http.ResponseWriter, r *http.Request) {}
