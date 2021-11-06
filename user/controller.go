package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const prefix = "/user"

type repository interface {
	getUser(id string) (User, error)
	getUsers() ([]UserItem, error)
}

type controller struct {
	repo repository
}

func InitializeController(repo repository, r *mux.Router) {
	c := controller{repo}
	s := r.PathPrefix(prefix).Subrouter()

	s.HandleFunc("/{id:[a-f0-9-]+}", c.getUser).Methods(http.MethodGet)
	s.HandleFunc("/all", c.getUsers).Methods(http.MethodGet)
}

func (c *controller) getUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := c.repo.getUser(id)
	if err != nil {
		http.Error(w, "User is not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *controller) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.repo.getUsers()
	if err != nil {
		http.Error(w, "Users is not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
