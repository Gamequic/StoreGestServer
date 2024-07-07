package users

import (
	"encoding/json"
	"net/http"
	"reflect"
	userservice "storegestserver/pkg/features/users/service"
	userstruct "storegestserver/pkg/features/users/struct"
	"storegestserver/utils"
	"storegestserver/utils/middlewares"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var logger *zap.Logger

// CRUD

func create(w http.ResponseWriter, r *http.Request) {
	logger = utils.NewLogger()
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

func find(w http.ResponseWriter, r *http.Request) {
	//Logger
	logger = utils.NewLogger()

	//Service
	var users []userservice.Users
	var httpsResponse int = userservice.Find(&users)

	//Https response
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(users)
}

func findOne(w http.ResponseWriter, r *http.Request) {
	logger = utils.NewLogger()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error("Failed to convert ID to integer", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user userservice.Users
	var httpsResponse int = userservice.FindOne(&user, uint(id))
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(user)
}

func update(w http.ResponseWriter, r *http.Request) {
	logger = utils.NewLogger()
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
	logger = utils.NewLogger()
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

	usersRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		middlewares.ValidatorHandler(http.HandlerFunc(create), reflect.TypeOf(userstruct.CreateUser{})).ServeHTTP(w, r)
	}).Methods("POST")
	usersRouter.HandleFunc("/", find).Methods("GET")
	usersRouter.HandleFunc("/{id}", findOne).Methods("GET")
	usersRouter.HandleFunc("/", update).Methods("PATCH")
	usersRouter.HandleFunc("/{id}", delete).Methods("DELETE")
}
