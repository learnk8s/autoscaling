package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
	log.Printf("Handling request from %v\n", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("> App listening on http://localhost:%v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic(err)
	}
}
