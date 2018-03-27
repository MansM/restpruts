package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	var events []Event
	if err := db.Find(&events).Error; err != nil {
		log.Fatal("error connecting to the database: ", err)
	} else {
		json.NewEncoder(w).Encode(events)
	}
}

// GetEvent Display just one event
func GetEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var event Event

	if i, err := strconv.Atoi(params["id"]); err == nil {
		db.Where("id = ?", i).First(&event)
	}
	requestedEvents.Inc()
	json.NewEncoder(w).Encode(event)
}

// CreateEvent blah
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		log.Panic("error in input:", err)
		w.WriteHeader(422) // unprocessable entity
		return
	}

	if err := db.Create(&event).Error; err != nil {
		log.Panic("error writing to db:", err)
		w.WriteHeader(500) // unprocessable entity
	} else {
		json.NewEncoder(w).Encode(event)
	}

	newEvents.Inc()
}

// DeleteEvent blah
func DeleteEvent(w http.ResponseWriter, r *http.Request) {}
