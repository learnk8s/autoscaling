package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

const port = 8080

var counter = promauto.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of HTTP requests processed.",
	},
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
	log.Printf("Handling request from %v\n", r.RemoteAddr)
	counter.Inc()
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", handler)

	fmt.Printf("> App listening on http://localhost:%v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic(err)
	}
}
