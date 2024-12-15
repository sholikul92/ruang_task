package model

import (
	"os"

	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
