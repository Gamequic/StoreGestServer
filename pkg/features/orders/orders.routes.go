package orders

import (
	"encoding/json"
	"net/http"
	"reflect"
	ordersservice "storegestserver/pkg/features/orders/service"
	ordersstruct "storegestserver/pkg/features/orders/struct"
	"storegestserver/utils/middlewares"
	"strconv"

	"github.com/gorilla/mux"
)

// var logger *zap.Logger

// CRUD

func create(w http.ResponseWriter, r *http.Request) {
	var Orders ordersservice.Orders

	/*
		This error is alredy been check it on middlewares.ValidatorHandler
		utils/middlewares/validatorHandler.go:29:68
	*/
	json.NewDecoder(r.Body).Decode(&Orders)

	ordersservice.Create(&Orders)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Orders)
}

func find(w http.ResponseWriter, r *http.Request) {
	//Service
	var Order []ordersservice.Orders
	var httpsResponse int = ordersservice.Find(&Order)

	//Https response
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(Order)
}

func findOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(middlewares.GormError{Code: 400, Message: err.Error(), IsGorm: true})
	}
	var Order ordersservice.Orders
	var httpsResponse int = ordersservice.FindOne(&Order, uint(id))
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(Order)
}

func findByDate(w http.ResponseWriter, r *http.Request) {
	// Date
	var Date ordersstruct.GetOrdersByDate
	json.NewDecoder(r.Body).Decode(&Date)
	// db
	var Orders []ordersservice.Orders
	var httpsResponse int = ordersservice.FindByDate(&Orders, Date)
	// response
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(Orders)
}

// Register function

func RegisterSubRoutes(router *mux.Router) {
	OrderRouter := router.PathPrefix("/orders").Subrouter()

	// Middlewares
	OrderRouter.Use(middlewares.AuthHandler)

	// ValidatorHandler
	orderCreateValidator := OrderRouter.NewRoute().Subrouter()
	orderCreateValidator.Use(middlewares.ValidatorHandler(reflect.TypeOf(ordersstruct.CreateOrders{})))
	orderUpdateValidator := OrderRouter.NewRoute().Subrouter()
	orderUpdateValidator.Use(middlewares.ValidatorHandler(reflect.TypeOf(ordersstruct.UpdateOrders{})))
	orderGetByDateValidator := OrderRouter.NewRoute().Subrouter()
	orderGetByDateValidator.Use(middlewares.ValidatorHandler(reflect.TypeOf(ordersstruct.GetOrdersByDate{})))

	orderCreateValidator.HandleFunc("/", create).Methods("POST")

	OrderRouter.HandleFunc("/", find).Methods("GET")
	OrderRouter.HandleFunc("/{id}", findOne).Methods("GET")
	orderGetByDateValidator.HandleFunc("/findByDate", findByDate).Methods("POST")
}
