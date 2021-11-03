package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there dick, I love %s!", r.URL.Path[1:])
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(os.Getenv("BASE_URL"), r))
}
