package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/users"

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

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", handler)
	users.HandleUsers(r)
	log.Fatal(http.ListenAndServe(os.Getenv("BASE_URL"), r))
}
