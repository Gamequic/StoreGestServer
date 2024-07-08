package auth

import (
	"encoding/json"
	"net/http"
	"reflect"
	authservice "storegestserver/pkg/features/auth/service"
	authstruct "storegestserver/pkg/features/auth/struct"
	"storegestserver/utils/middlewares"

	"github.com/gorilla/mux"
)

// CRUD

func LogIn(w http.ResponseWriter, r *http.Request) {
	var user authstruct.LogIn

	json.NewDecoder(r.Body).Decode(&user)

	var token string = authservice.LogIn(&user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}

// func ChangePassword(w http.ResponseWriter, r *http.Request) {
// 	//Service
// 	var users []userservice.Users
// 	var httpsResponse int = userservice.Find(&users)

// 	//Https response
// 	w.WriteHeader(httpsResponse)
// 	json.NewEncoder(w).Encode(users)
// }

// func RequestPasswordChange(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		panic(middlewares.GormError{Code: 400, Message: err.Error(), IsGorm: true})
// 	}
// 	var user userservice.Users
// 	var httpsResponse int = userservice.FindOne(&user, uint(id))
// 	w.WriteHeader(httpsResponse)
// 	json.NewEncoder(w).Encode(user)
// }

// func ApplyPasswordChange(w http.ResponseWriter, r *http.Request) {
// 	var user userservice.Users
// 	json.NewDecoder(r.Body).Decode(&user)
// 	userservice.Update(&user)
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(user)
// }

// Register function

func RegisterSubRoutes(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()

	// ValidatorHandler
	authValidator := authRouter.NewRoute().Subrouter()
	authValidator.Use(middlewares.ValidatorHandler(reflect.TypeOf(authstruct.LogIn{})))

	authValidator.HandleFunc("/login", LogIn).Methods("POST")

	// authRouter.HandleFunc("/changePassword", ChangePassword).Methods("GET")
	// authRouter.HandleFunc("/RequestPasswordChange", RequestPasswordChange).Methods("DELETE")
	// authRouter.HandleFunc("/ApplyPasswordChange", ApplyPasswordChange).Methods("DELETE")
}
