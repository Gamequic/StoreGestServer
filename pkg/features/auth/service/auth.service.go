package authservice

import (
	"fmt"
	"os"
	authstruct "storegestserver/pkg/features/auth/struct"
	userservice "storegestserver/pkg/features/users/service"
	"storegestserver/utils/middlewares"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var Logger *zap.Logger

// Initialize the auth service
func InitAuthService() {

}

// Auth Operations

func LogIn(u *authstruct.LogIn) string {
	var jwtKey = []byte(os.Getenv("JWTSECRET"))

	// Check if user exists
	var user userservice.Users
	userservice.FindByEmail(&user, u.Email)

	//Check if password is correct
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			panic(middlewares.GormError{Code: 401, Message: "Password is wrong", IsGorm: true})
		} else {
			panic(err.Error())
		}
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	TokenData := &authstruct.TokenStruct{
		Username: user.Name,
		Email:    user.Email,
		Id:       int(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenData)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

// func RequestPasswordChange(user *Users, id uint) int {
// 	if err := database.DB.First(user, id).Error; err != nil {
// 		if err.Error() == "record not found" {
// 			panic(middlewares.GormError{Code: 404, Message: "Users not found", IsGorm: true})
// 		} else {
// 			panic(err)
// 		}
// 	}
// 	return http.StatusOK
// }

// func ApplyPasswordChange(u *Users) {
// 	// No autorize editing no existing users
// 	var previousUsers Users
// 	FindOne(&previousUsers, uint(u.ID))

// 	// Encrypt password
// 	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
// 	u.Password = string(bytes)

// 	if err := database.DB.Save(u).Error; err != nil {
// 		if err.Error() == `ERROR: duplicate key value violates unique constraint "uni_users_email" (SQLSTATE 23505)` {
// 			panic(middlewares.GormError{Code: 409, Message: "Email is on use", IsGorm: true})
// 		} else {
// 			panic(err)
// 		}
// 	}
// }

// func Delete(id int) {
// 	Logger = utils.NewLogger()

// 	// No autorize deleting no existing users
// 	var previousUsers Users
// 	FindOne(&previousUsers, uint(id))

// 	if err := database.DB.Delete(&Users{}, id).Error; err != nil {
// 		panic(err)
// 	}
// }
