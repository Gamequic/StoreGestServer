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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

// func RequestPasswordChange(w http.ResponseWriter, r *http.Request) {
// 	var body authstruct.RequestChangePassword
// 	json.NewDecoder(r.Body).Decode(&body)
// 	var response map[string]interface{} = authservice.RequestPasswordChange(body.Email)
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
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
	authLogInValidator := authRouter.NewRoute().Subrouter()
	authLogInValidator.Use(middlewares.ValidatorHandler(reflect.TypeOf(authstruct.LogIn{})))

	// authRequestPassword := authRouter.NewRoute().Subrouter()
	// authRequestPassword.Use(middlewares.ValidatorHandler(reflect.TypeOf(authstruct.RequestChangePassword{})))

	// Endpoints
	authLogInValidator.HandleFunc("/login", LogIn).Methods("POST")
	// authRequestPassword.HandleFunc("/requestchangepassword", RequestPasswordChange).Methods("POST")
}
