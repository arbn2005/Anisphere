package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(Password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), err
}

func CheckPassword(HashPassword string, Password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashPassword), []byte(Password))

	if err == nil {
		return true
	} else {
		return false
	}
}
