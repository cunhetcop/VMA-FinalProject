// utils/password.go
package utils

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var customRand *rand.Rand

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	customRand = rand.New(s)
}

func GenerateRandomPassword(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = charset[customRand.Intn(len(charset))]
	}
	return string(password)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}