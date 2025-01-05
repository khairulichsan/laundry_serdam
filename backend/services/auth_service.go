package services

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID uint) (string, error) {
	// Create the Claims
	claims := &jwt.StandardClaims{
		Subject:   strconv.Itoa(int(userID)),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
