package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there dick, I love %s!", r.URL.Path[1:])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
