package main

// Event type
type Event struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var events []Event
