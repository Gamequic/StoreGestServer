package food

import (
	"encoding/json"
	"net/http"
	"reflect"
	foodservice "storegestserver/pkg/features/food/service"
	foodstruct "storegestserver/pkg/features/food/struct"
	"storegestserver/utils"
	"storegestserver/utils/middlewares"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var logger *zap.Logger

// CRUD

func create(w http.ResponseWriter, r *http.Request) {
	var food foodservice.Food

	/*
		This error is alredy been check it on middlewares.ValidatorHandler
		utils/middlewares/validatorHandler.go:29:68
	*/
	json.NewDecoder(r.Body).Decode(&food)

	foodservice.Create(&food)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(food)
}

func find(w http.ResponseWriter, r *http.Request) {
	//Service
	var foods []foodservice.Food
	var httpsResponse int = foodservice.Find(&foods)

	//Https response
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(foods)
}

func findOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(middlewares.GormError{Code: 400, Message: err.Error(), IsGorm: true})
	}
	var food foodservice.Food
	var httpsResponse int = foodservice.FindOne(&food, uint(id))
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(food)
}

func update(w http.ResponseWriter, r *http.Request) {
	var food foodservice.Food
	json.NewDecoder(r.Body).Decode(&food)

	foodservice.Update(&food)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(food)
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
	foodservice.Delete(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Food deleted successfully")
}

// Register function

func RegisterSubRoutes(router *mux.Router) {
	foodRouter := router.PathPrefix("/food").Subrouter()

	// Middlewares
	foodRouter.Use(middlewares.AuthHandler)

	// ValidatorHandler
	foodCreateValidator := foodRouter.NewRoute().Subrouter()
	foodCreateValidator.Use(middlewares.ValidatorHandler(reflect.TypeOf(foodstruct.CreateFood{})))

	foodUpdateValidator := foodRouter.NewRoute().Subrouter()
	foodUpdateValidator.Use(middlewares.ValidatorHandler(reflect.TypeOf(foodstruct.UpdateFood{})))

	foodCreateValidator.HandleFunc("/", create).Methods("POST")
	foodUpdateValidator.HandleFunc("/", update).Methods("PATCH")

	foodRouter.HandleFunc("/", find).Methods("GET")
	foodRouter.HandleFunc("/{id}", findOne).Methods("GET")
	foodRouter.HandleFunc("/{id}", delete).Methods("DELETE")
}
