package moneyservice

import (
	"errors"
	"fmt"
	"net/http"
	"storegestserver/pkg/database"
	moneystruct "storegestserver/pkg/features/money/struct"
	"storegestserver/utils"
	"storegestserver/utils/middlewares"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Logger *zap.Logger

// Define the user model
type Money struct {
	gorm.Model         // Embedding gorm.Model for default fields like ID, CreatedAt, UpdatedAt, DeletedAt
	Amount      int    `gorm:"not null"`
	Current     int    `gorm:"not null"`
	Reason      string `gorm:"omitempty"`
	Description string `gorm:"not null"`
}

// Initialize the money service
func InitMoneyService() {
	Logger = utils.NewLogger()
	err := database.DB.AutoMigrate(&Money{})
	if err != nil {
		Logger.Error(fmt.Sprint("Failed to migrate database:", err))
	}
}

// CRUD Operations

func Create(M *Money) {
	var existingMoney Money

	if err := database.DB.Order("created_at desc").First(&existingMoney).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			panic(err)
		}
	}

	if existingMoney.ID != 0 {
		M.Current = existingMoney.Current + M.Amount
	} else {
		M.Current = M.Amount
	}

	if err := database.DB.Create(M).Error; err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "uni_money_date"`) {
			panic(middlewares.GormError{Code: http.StatusConflict, Message: "Date is already in use", IsGorm: true})
		} else {
			panic(err)
		}
	}
}

func Find(u *[]Money) int {
	database.DB.Find(u)
	return http.StatusOK
}

func FindOne(Money *Money, id uint) int {
	if err := database.DB.First(Money, id).Error; err != nil {
		if err.Error() == "record not found" {
			panic(middlewares.GormError{Code: 404, Message: "Record not found", IsGorm: true})
		} else {
			panic(err)
		}
	}
	return http.StatusOK
}

func FindByDate(moneyRecords *[]Money, body moneystruct.GetMoneyByDate) int {

	startDate := time.Date(int(body.Year), time.Month(body.Month), int(body.Day), 0, 0, 0, 0, time.UTC)
	endDate := time.Date(int(body.Year), time.Month(body.Month), int(body.Day)+1, 0, 0, 0, 0, time.UTC)

	database.DB.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&moneyRecords)

	return http.StatusOK
}
