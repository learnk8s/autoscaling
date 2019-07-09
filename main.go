package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("> App listening on http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
