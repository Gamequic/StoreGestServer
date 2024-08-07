package money

import (
	"encoding/json"
	"net/http"
	"reflect"
	moneyservice "storegestserver/pkg/features/money/service"
	moneystruct "storegestserver/pkg/features/money/struct"
	"storegestserver/utils/middlewares"
	"strconv"

	"github.com/gorilla/mux"
)

// var logger *zap.Logger

// CRUD

func create(w http.ResponseWriter, r *http.Request) {
	var user moneyservice.Money

	/*
		This error is alredy been check it on middlewares.ValidatorHandler
		utils/middlewares/validatorHandler.go:29:68
	*/
	json.NewDecoder(r.Body).Decode(&user)

	moneyservice.Create(&user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func find(w http.ResponseWriter, r *http.Request) {
	//Service
	var Money []moneyservice.Money
	var httpsResponse int = moneyservice.Find(&Money)

	//Https response
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(Money)
}

func findLastOne(w http.ResponseWriter, r *http.Request) {
	var Money moneyservice.Money
	var httpsResponse int = moneyservice.FindLastOne(&Money)
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(Money)
}

func findOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(middlewares.GormError{Code: 400, Message: err.Error(), IsGorm: true})
	}
	var Money moneyservice.Money
	var httpsResponse int = moneyservice.FindOne(&Money, uint(id))
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(Money)
}

func findByDate(w http.ResponseWriter, r *http.Request) {
	// Date
	var Date moneystruct.GetMoneyByDate
	json.NewDecoder(r.Body).Decode(&Date)
	// db
	var Money []moneyservice.Money
	var httpsResponse int = moneyservice.FindByDate(&Money, Date)
	// response
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(Money)
}

func findByDateRange(w http.ResponseWriter, r *http.Request) {
	// Date
	var Date moneystruct.GetMoneyByDateRange
	json.NewDecoder(r.Body).Decode(&Date)
	// db
	var Money []moneyservice.Money
	var httpsResponse int = moneyservice.FindByDateRange(&Money, Date)
	// response
	w.WriteHeader(httpsResponse)
	json.NewEncoder(w).Encode(Money)
}

// Register function

func RegisterSubRoutes(router *mux.Router) {
	moneyRouter := router.PathPrefix("/money").Subrouter()

	// ValidatorHandler
	MoneyCreatorValidtor := moneyRouter.NewRoute().Subrouter()
	MoneyCreatorValidtor.Use(middlewares.ValidatorHandler(reflect.TypeOf(moneystruct.CreateMoney{})))
	MoneyGetByDate := moneyRouter.NewRoute().Subrouter()
	MoneyGetByDate.Use(middlewares.ValidatorHandler(reflect.TypeOf(moneystruct.GetMoneyByDate{})))
	MoneyGetByDateRange := moneyRouter.NewRoute().Subrouter()
	MoneyGetByDateRange.Use(middlewares.ValidatorHandler(reflect.TypeOf(moneystruct.GetMoneyByDateRange{})))
	moneyRouter.Use(middlewares.AuthHandler)

	MoneyCreatorValidtor.HandleFunc("/", create).Methods("POST")
	moneyRouter.HandleFunc("/", find).Methods("GET")
	moneyRouter.HandleFunc("/lastOne", findLastOne).Methods("GET")
	moneyRouter.HandleFunc("/{id}", findOne).Methods("GET")
	MoneyGetByDate.HandleFunc("/findByDate", findByDate).Methods("POST")
	MoneyGetByDateRange.HandleFunc("/findByDateRange", findByDateRange).Methods("POST")
}
