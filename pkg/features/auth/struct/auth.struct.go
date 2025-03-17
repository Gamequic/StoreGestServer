package authstruct

import "github.com/golang-jwt/jwt"

type TokenStruct struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email    string
	Id       int
}

type LogIn struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type UserData struct {
	Email    string `validate:"required,email"`
	Name     string `validate:"required"`
	Password string `validate:"required,min=8"`
	Token    string `validate:"required,jwt"`
}

// type RequestChangePassword struct {
// 	Email string `validate:"required,email"`
// }
