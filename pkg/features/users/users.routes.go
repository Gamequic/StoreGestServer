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
	var user userservice.Users

	/*
		This error is alredy been check it on middlewares.ValidatorHandler
		utils/middlewares/validatorHandler.go:29:68
	*/
	json.NewDecoder(r.Body).Decode(&user)

	userservice.Create(&user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func find(w http.ResponseWriter, r *http.Request) {
	//Service
	var users []userservice.Users
	var httpsResponse int = userservice.Find(&users)

	//Https response
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(users)
}

func findOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(middlewares.GormError{Code: 400, Message: err.Error(), IsGorm: true})
	}
	var user userservice.Users
	var httpsResponse int = userservice.FindOne(&user, uint(id))
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(user)
}

func update(w http.ResponseWriter, r *http.Request) {
	var user userservice.Users
	json.NewDecoder(r.Body).Decode(&user)
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

	// ValidatorHandler
	usersUpdateValidtor := usersRouter.NewRoute().Subrouter()
	usersUpdateValidtor.Use(middlewares.ValidatorHandler(reflect.TypeOf(userstruct.UpdateUser{})))

	userCreateValidator := usersRouter.NewRoute().Subrouter()
	userCreateValidator.Use(middlewares.ValidatorHandler(reflect.TypeOf(userstruct.CreateUser{})))

	usersUpdateValidtor.HandleFunc("/", update).Methods("PATCH")
	userCreateValidator.HandleFunc("/", create).Methods("POST")

	usersRouter.HandleFunc("/", find).Methods("GET")
	usersRouter.HandleFunc("/{id}", findOne).Methods("GET")
	usersRouter.HandleFunc("/{id}", delete).Methods("DELETE")
}
