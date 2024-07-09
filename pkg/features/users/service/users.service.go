package userservice

import (
	"fmt"
	"net/http"
	"storegestserver/pkg/database"
	"storegestserver/utils"
	"storegestserver/utils/middlewares"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var Logger *zap.Logger

// Define the user model
type Users struct {
	gorm.Model        // Embedding gorm.Model for default fields like ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"not null"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
}

// Initialize the user service
func InitUsersService() {
	Logger = utils.NewLogger()
	err := database.DB.AutoMigrate(&Users{})
	if err != nil {
		Logger.Error(fmt.Sprint("Failed to migrate database:", err))
	}
}

// CRUD Operations

func Create(u *Users) {
	// Encrypt password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)

	if err := database.DB.Create(u).Error; err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "uni_users_email" (SQLSTATE 23505)` {
			panic(middlewares.GormError{Code: 409, Message: "Email is on use", IsGorm: true})
		} else {
			panic(err)
		}
	}
}

func Find(u *[]Users) int {
	database.DB.Find(u)
	return http.StatusOK
}

func FindOne(user *Users, id uint) int {
	if err := database.DB.First(user, id).Error; err != nil {
		if err.Error() == "record not found" {
			panic(middlewares.GormError{Code: 404, Message: "Users not found", IsGorm: true})
		} else {
			panic(err)
		}
	}
	return http.StatusOK
}

func FindByEmail(user *Users, email string) int {
	if err := database.DB.Where("email = ?", email).First(user).Error; err != nil {
		if err.Error() == "record not found" {
			panic(middlewares.GormError{Code: 404, Message: "User not found", IsGorm: true})
		} else {
			panic(err)
		}
	}
	return http.StatusOK
}

func Update(u *Users, userId uint) {
	// No autorize editing no existing users
	var previousUsers Users
	FindOne(&previousUsers, uint(u.ID))

	// Is the same user?
	if u.ID != userId {
		panic(middlewares.GormError{Code: http.StatusNotAcceptable, Message: "Is not allow to modify others users", IsGorm: true})
	}

	// Encrypt password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)

	if err := database.DB.Save(u).Error; err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "uni_users_email" (SQLSTATE 23505)` {
			panic(middlewares.GormError{Code: 409, Message: "Email is on use", IsGorm: true})
		} else {
			panic(err)
		}
	}
}

func Delete(id int) {
	Logger = utils.NewLogger()

	// No autorize deleting no existing users
	var previousUsers Users
	FindOne(&previousUsers, uint(id))

	if err := database.DB.Delete(&Users{}, id).Error; err != nil {
		panic(err)
	}
}
