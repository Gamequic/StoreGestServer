package foodservice

import (
	"fmt"
	"net/http"
	"storegestserver/pkg/database"
	"storegestserver/utils"
	"storegestserver/utils/middlewares"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Logger *zap.Logger

// Define the food model
type Food struct {
	gorm.Model        // Embedding gorm.Model for default fields like ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"not null; unique"`
	Amount     uint   `gorm:"not null"`
	Photo      string `gorm:"not null"`
	IsKg       bool   `gorm:"not null"`
}

// Initialize the user service
func InitFoodService() {
	Logger = utils.NewLogger()
	err := database.DB.AutoMigrate(&Food{})
	if err != nil {
		Logger.Error(fmt.Sprint("Failed to migrate database:", err))
	}
}

// CRUD Operations

func Create(u *Food) {
	if err := database.DB.Create(u).Error; err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "uni_foods_name" (SQLSTATE 23505)` {
			panic(middlewares.GormError{Code: 409, Message: "Name is on use", IsGorm: true})
		} else {
			panic(err)
		}
	}
}

func Find(u *[]Food) int {
	database.DB.Find(u)
	return http.StatusOK
}

func FindOne(Food *Food, id uint) int {
	if err := database.DB.First(Food, id).Error; err != nil {
		if err.Error() == "record not found" {
			panic(middlewares.GormError{Code: 404, Message: "Food not found", IsGorm: true})
		} else {
			panic(err)
		}
	}
	return http.StatusOK
}

func Update(u *Food) {
	// No autorize editing no existing food
	var previousFood Food
	FindOne(&previousFood, uint(u.ID))

	if err := database.DB.Save(u).Error; err != nil {
		panic(err)
	}
}

func Delete(id int) {
	// No autorize editing no existing food
	var previousFood Food
	FindOne(&previousFood, uint(id))

	if err := database.DB.Delete(&Food{}, id).Error; err != nil {
		panic(err)
	}
}
