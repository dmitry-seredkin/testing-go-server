package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const prefix = "/user"

type repository interface {
	createUser(user CreateUser) (User, error)
	deleteUser(id string) error
	getUser(id string) (User, error)
	getUsers() ([]UserItem, error)
	loginUser(user LoginUser) (bool, error)
}

type controller struct {
	repo repository
}

func InitializeController(repo repository, r *mux.Router) {
	c := controller{repo}
	s := r.PathPrefix(prefix).Subrouter()

	s.HandleFunc("/all", c.getUsers).Methods(http.MethodGet)
	s.HandleFunc("/new", c.createUser).Methods(http.MethodPost)
	s.HandleFunc("/login", c.loginUser).Methods(http.MethodPost)
	s.HandleFunc("/{id:[a-f0-9-]+}", c.getUser).Methods(http.MethodGet)
	s.HandleFunc("/{id:[a-f0-9-]+}", c.deleteUser).Methods(http.MethodDelete)
}

func (c *controller) createUser(w http.ResponseWriter, r *http.Request) {
	var createUser CreateUser

	decodeErr := json.NewDecoder(r.Body).Decode(&createUser)
	if decodeErr != nil {
		http.Error(w, "Can't decode user information", http.StatusUnprocessableEntity)
		return
	}

	user, err := c.repo.createUser(createUser)
	if err != nil {
		http.Error(w, "User creation fail", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *controller) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := c.repo.deleteUser(id)
	if err != nil {
		http.Error(w, "User is not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("ok"))
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

func (c *controller) loginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser LoginUser
	decodeErr := json.NewDecoder(r.Body).Decode(&loginUser)
	if decodeErr != nil {
		http.Error(w, "Can't decode user information", http.StatusUnprocessableEntity)
		return
	}

	isValid, err := c.repo.loginUser(loginUser)
	if err != nil {
		http.Error(w, "Users is not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(isValid)
}
