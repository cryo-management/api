package models

import jwt "github.com/dgrijalva/jwt-go"

// User defines the struct of this object
type User struct {
	ID           string `json:"id" sql:"id" pk:"true"`
	FirstName    string `json:"first_name" sql:"first_name"`
	LastName     string `json:"last_name" sql:"last_name"`
	Email        string `json:"email" sql:"email"`
	Password     string `json:"password" sql:"password"`
	LanguageCode string `json:"language" sql:"language"`
	Active       bool   `json:"active" sql:"active"`
}

// UserCustomClaims used to parse token payload
type UserCustomClaims struct {
	User User
	jwt.StandardClaims
}
