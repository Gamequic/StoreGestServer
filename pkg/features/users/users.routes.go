package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

// CRUD

func createUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create users"))
}

func readUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Read users"))
}

func updateUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update users"))
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete users"))
}

// Registrer function

func RegisterSubRoutes(router *mux.Router) {
	usersRouter := router.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("/", createUsers).Methods("POST")
	usersRouter.HandleFunc("/", readUsers).Methods("GET")
	usersRouter.HandleFunc("/", updateUsers).Methods("PATCH")
	usersRouter.HandleFunc("/", deleteUsers).Methods("DELETE")

}
