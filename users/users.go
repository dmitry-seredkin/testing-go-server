package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var users = []User{}

func HandleUsers(router *mux.Router) {
	subRouter := router.PathPrefix("/users").Subrouter()
	subRouter.HandleFunc("", returnUsers).Methods("GET")
	subRouter.HandleFunc("", createUser).Methods("POST")
	subRouter.HandleFunc("/{id:[0-9]+}", returnUser).Methods("GET")
	subRouter.HandleFunc("/{id:[0-9]+}", updateUser).Methods("PUT")
	subRouter.HandleFunc("/{id:[0-9]+}", deleteUser).Methods("DELETE")
}

func returnUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal("Couldn't decode request body")
	}

	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func returnUser(w http.ResponseWriter, r *http.Request) {
	id := getIdFromVars(r)

	for _, user := range users {
		if user.Id == id {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := getIdFromVars(r)

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		log.Fatal("Couldn't decode request body")
	}

	for i, user := range users {
		if user.Id == id {
			users[i] = newUser
			json.NewEncoder(w).Encode(newUser)
			return
		}
	}

	http.Error(w, "User with this id does not exist", http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := getIdFromVars(r)

	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)
			w.Write([]byte("ok"))
			return
		}
	}

	http.Error(w, "User with this id does not exist", http.StatusNotFound)
}

// HELPERS
func getIdFromVars(r *http.Request) string {
	id := mux.Vars(r)["id"]

	return id
}
