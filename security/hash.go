package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(data string) string {
	pass, _ := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	return string(pass)
}

func CheckPassword(encripted string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encripted), []byte(password))
	if err == nil {
		return true
	} else {
		return false
	}
}
