package ordersservice

import (
	"fmt"
	"net/http"
	"sort"
	"storegestserver/pkg/database"
	foodservice "storegestserver/pkg/features/food/service"
	moneyservice "storegestserver/pkg/features/money/service"
	ordersstruct "storegestserver/pkg/features/orders/struct"
	"storegestserver/utils"
	"storegestserver/utils/middlewares"
	"strings"
	"time"

	"github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Logger *zap.Logger

// Define the order model
type Orders struct {
	gorm.Model                 // Embedding gorm.Model for default fields like ID, CreatedAt, UpdatedAt, DeletedAt
	FoodList   pq.Int64Array   `gorm:"type:integer[]"`
	FoodAmount pq.Float64Array `gorm:"type:float[]"`
	Amount     float64         `gorm:"not null"`
	MoneyId    uint            `gorm:"not null"`
}

// Initialize the money service
func InitOrdersService() {
	Logger = utils.NewLogger()
	err := database.DB.AutoMigrate(&Orders{})
	if err != nil {
		Logger.Error(fmt.Sprint("Failed to migrate database:", err))
	}
}

// CRUD Operations

func Create(Order *Orders) {

	// Check if it is the same lenght of two list
	if len(Order.FoodList) != len(Order.FoodAmount) {
		panic(middlewares.GormError{Code: http.StatusBadRequest, Message: "Error in food list and amount, Line 46 order.service.go", IsGorm: true})
	}

	// Check if all items exist
	var food foodservice.Food
	for _, foodID := range Order.FoodList {
		food = foodservice.Food{}
		foodservice.FindOne(&food, uint(foodID))
	}

	// Add money operation to db
	var MoneyOperation moneyservice.Money
	MoneyOperation.Amount = float64(Order.Amount)
	MoneyOperation.Reason = "Compra"
	moneyservice.Create(&MoneyOperation)

	// Add MoneyId to Order
	Order.MoneyId = MoneyOperation.ID

	// Convert FoodList to pq.Int64Array
	Order.FoodList = pq.Int64Array(Order.FoodList)

	// Add order to db
	if err := database.DB.Create(Order).Error; err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "uni_money_date"`) {
			panic(middlewares.GormError{Code: http.StatusConflict, Message: "Date is already in use", IsGorm: true})
		} else {
			panic(err)
		}
	}
}
func Find(u *[]Orders) int {
	database.DB.Find(u)
	return http.StatusOK
}

func FindOne(Orders *Orders, id uint) int {
	if err := database.DB.First(Orders, id).Error; err != nil {
		if err.Error() == "record not found" {
			panic(middlewares.GormError{Code: 404, Message: "Record not found", IsGorm: true})
		} else {
			panic(err)
		}
	}
	return http.StatusOK
}

func FindByDate(OrdersRecord *[]Orders, body ordersstruct.GetOrdersByDate) int {
	// Set the time for the correct timezone
	// To-do
	// [ ] Load this from .env
	location, err := time.LoadLocation("America/Mexico_City")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return http.StatusInternalServerError
	}

	startDate := time.Date(int(body.Year), time.Month(body.Month), int(body.Day), 0, 0, 0, 0, location)
	endDate := time.Date(int(body.Year), time.Month(body.Month), int(body.Day), 23, 59, 59, 999999999, location)

	database.DB.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&OrdersRecord)

	return http.StatusOK
}

func FindByDateRange(ordersRecords *[]Orders, body ordersstruct.GetOrdersByDateRange) int {

	// Set the time for the correct timezone
	// To-do
	// [ ] Load this from .env
	location, err := time.LoadLocation("America/Mexico_City")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return http.StatusInternalServerError
	}

	startDate := time.Date(int(body.InitYear), time.Month(body.InitMonth), int(body.InitDay), 0, 0, 0, 0, location)
	endDate := time.Date(int(body.EndYear), time.Month(body.EndMonth), int(body.EndDay), 23, 59, 59, 999999999, location)

	database.DB.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(ordersRecords)

	return http.StatusOK
}

func Statistics(Statistics *ordersstruct.Statistics, body ordersstruct.GetOrdersByDateRange) int {
	var Orders []Orders

	FindByDateRange(&Orders, body)

	var average float64 = 0
	var ordersNumber int = 0

	salesPerItem := make(map[int]float64)

	// get a list of the amount of sales of each product
	for _, order := range Orders {
		for index, foodId := range order.FoodList {
			_, exists := salesPerItem[int(foodId)]
			if exists {
				salesPerItem[int(foodId)] = salesPerItem[int(foodId)] + float64(order.FoodAmount[index])
			} else {
				salesPerItem[int(foodId)] = float64(order.FoodAmount[index])
			}
		}
		average = average + float64(order.Amount)
		ordersNumber++
	}

	/* order the list from largest to smallest */
	// Step 1: Get the keys from the map
	keys := make([]int, 0, len(salesPerItem))
	for k := range salesPerItem {
		keys = append(keys, k)
	}
	// Step 2: Sort the keys based on the map values
	sort.Slice(keys, func(i, j int) bool {
		return salesPerItem[keys[i]] > salesPerItem[keys[j]]
	})

	// Get the first 3 products and added to the petition
	for i := 0; i < 3 && i < len(keys); i++ {
		var food foodservice.Food
		foodservice.FindOne(&food, uint(keys[i]))
		Statistics.Products = append(Statistics.Products, food)
	}

	average = average / float64(len(Orders))

	Statistics.Average = average
	Statistics.OrdersNumber = ordersNumber

	return http.StatusOK
}
