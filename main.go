package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/prometheus/client_golang/prometheus"
)

var db = NewDB()

func main() {
	prometheus.MustRegister(newEvents)
	prometheus.MustRegister(requestedEvents)

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
