package userservice

import (
	"fmt"
	"storegestserver/pkg/database"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Logger *zap.Logger

// Define the user model
type Users struct {
	gorm.Model // Embedding gorm.Model for default fields like ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string
	Email      string `gorm:"unique"`
	Password   string
}

// Initialize the user service
func InitUsersService() {
	err := database.DB.AutoMigrate(&Users{})
	if err != nil {
		Logger.Error(fmt.Sprint("Failed to migrate database:", err))
	}
}

// CRUD Operations

func Create(u *Users) {
	if err := database.DB.Create(u).Error; err != nil {
		Logger.Error("Failed to create user:", zap.Error(err))
	}
}

func Find(users *[]Users) {
	if err := database.DB.Find(users).Error; err != nil {
		Logger.Error("Failed to find users:", zap.Error(err))
	}
}

func FindOne(user *Users, id uint) {
	if err := database.DB.First(user, id).Error; err != nil {
		Logger.Error("Failed to find user:", zap.Error(err))
	}
}

func Update(u *Users) {
	if err := database.DB.Save(u).Error; err != nil {
		Logger.Error("Failed to update user:", zap.Error(err))
	}
}

func Delete(id int) {
	if err := database.DB.Delete(&Users{}, id).Error; err != nil {
		Logger.Error("Failed to delete user:", zap.Error(err))
	}
}
