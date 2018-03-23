package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := NewRouter()
	router.HandleFunc("/event", GetEvents).Methods("GET")
	router.HandleFunc("/event/{id}", GetEvent).Methods("GET")
	return router
}

func TestGetEvents(t *testing.T) {
	request, _ := http.NewRequest("GET", "/event", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestGetEvent(t *testing.T) {
	request, _ := http.NewRequest("GET", "/event/1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
