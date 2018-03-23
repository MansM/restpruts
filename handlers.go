package main

import (
	"encoding/json"
	"fmt"
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
	json.NewEncoder(w).Encode(&Event{})
}

// CreateEvent blah
func CreateEvent(w http.ResponseWriter, r *http.Request) {}

// DeleteEvent blah
func DeleteEvent(w http.ResponseWriter, r *http.Request) {}
