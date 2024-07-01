package users

import (
	"encoding/json"
	"net/http"
	userservice "storegestserver/pkg/features/users/service"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var logger *zap.Logger

// CRUD

func create(w http.ResponseWriter, r *http.Request) {
	var user userservice.Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Error("Failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userservice.Create(&user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func read(w http.ResponseWriter, r *http.Request) {
	var users []userservice.Users
	userservice.Find(&users)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func findOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error("Failed to convert ID to integer", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user userservice.Users
	userservice.FindOne(&user, uint(id))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func update(w http.ResponseWriter, r *http.Request) {
	var user userservice.Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Error("Failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userservice.Update(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error("Failed to convert ID to integer", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userservice.Delete(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User deleted successfully")
}

// Register function

func RegisterSubRoutes(router *mux.Router) {
	usersRouter := router.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("/", create).Methods("POST")
	usersRouter.HandleFunc("/", read).Methods("GET")
	usersRouter.HandleFunc("/{id}", findOne).Methods("GET")
	usersRouter.HandleFunc("/", update).Methods("PATCH")
	usersRouter.HandleFunc("/{id}", delete).Methods("DELETE")
}
