package controllers

import (
	"example.com/enterprise-grade-system/pkg/models"
	"example.com/enterprise-grade-system/pkg/services"
	"encoding/json"
	"net/http"
)

// GetUsers returns a list of users
:func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// GetUser returns a single user
:func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user, err := services.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}