package ordersservice

import (
	"fmt"
	"net/http"
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
	gorm.Model               // Embedding gorm.Model for default fields like ID, CreatedAt, UpdatedAt, DeletedAt
	FoodList   pq.Int64Array `gorm:"type:integer[]"`
	Amount     uint          `gorm:"not null"`
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

	// Check if all items exist
	var food foodservice.Food
	for _, foodID := range Order.FoodList {
		foodservice.FindOne(&food, uint(foodID))
	}

	// Add money operation to db
	var MoneyOperation moneyservice.Money
	MoneyOperation.Amount = int(Order.Amount)
	MoneyOperation.Reason = "Compra"
	moneyservice.Create(&MoneyOperation)

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

	startDate := time.Date(int(body.Year), time.Month(body.Month), int(body.Day), 0, 0, 0, 0, time.UTC)
	endDate := time.Date(int(body.Year), time.Month(body.Month), int(body.Day)+1, 0, 0, 0, 0, time.UTC)

	database.DB.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&OrdersRecord)

	return http.StatusOK
}