package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("anisphere-secrect") // A temp password must be loded from env later on

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int64) (string, error) {
	register_claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	claims := Claims{
		UserID:           int(userID),
		RegisteredClaims: register_claims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
