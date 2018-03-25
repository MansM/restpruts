package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	newEvents = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "events_new_total",
		Help: "new events",
	})
	requestedEvents = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "events_requested",
		Help: "requested events",
	})
)
