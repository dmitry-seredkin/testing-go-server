package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const prefix = "/user"

type repository interface {
	getUser(id string) (User, error)
}

type controller struct {
	repo repository
}

func InitializeController(repo repository, r *mux.Router) {
	c := controller{repo}
	s := r.PathPrefix(prefix).Subrouter()

	s.HandleFunc("/{id:[a-f0-9-]+}", c.getUser).Methods(http.MethodGet)
}

func (c *controller) getUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := c.repo.getUser(id)
	if (err != nil) {
		http.Error(w, "User is not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
