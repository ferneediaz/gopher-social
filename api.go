package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

var users = []User{}

func (a *api) usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.getUsersHandler(w, r)
	case http.MethodPost:
		a.createUsersHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Encode users slice to JSON
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// No need to set status here; default is 200 OK
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request body to User struct
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate user input before inserting
	if err = insertUser(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add the user to the users slice
	users = append(users, payload)

	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	// Input validation
	if u.FirstName == "" {
		return errors.New("first name is required")
	}
	if u.LastName == "" {
		return errors.New("last name is required")
	}

	// Check for duplicates
	for _, user := range users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return errors.New("user already exists")
		}
	}
	return nil
}
