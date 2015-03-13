package services

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	passwordBytes := []byte(password)

	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordBytes, 10)
	return string(hashedPassword)
}

func ComparePassword(password string, hash string) bool {
	passwordBytes := []byte(password)
	hashBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)
	if err == nil {
		return true
	} else {
		return false
	}
}
