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
