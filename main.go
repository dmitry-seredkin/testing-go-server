package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/user"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := pgx.Connect(context.Background(), os.ExpandEnv(os.Getenv("DB_URL")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	router := mux.NewRouter().StrictSlash(true)
	APIRouter := router.PathPrefix("/api/v1").Subrouter()
	router.Handle("/", http.FileServer(http.Dir("./views/")))

	userRepository := user.InitializeRepository(conn)
	user.InitializeController(userRepository, APIRouter)

	log.Fatal(http.ListenAndServe(os.Getenv("BASE_URL"), router))
}
