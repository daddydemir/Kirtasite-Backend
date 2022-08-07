package auth

import (
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/daddydemir/kirtasiye-projesi/service"
)

func FileAuthUser(tokenString string, userId int) (bool, map[string]interface{}) {
	tokenUser, status := repositories.GetCustomerByUserName(TokenParser(tokenString))
	var customer models.Customer
	customer = tokenUser.(models.Customer)
	if status {
		if customer.UserId == userId {
			return true, service.OkMessage()
		} else {
			return false, service.YetkisiOlmayanMessage()
		}
	} else {
		return false, service.YetkisiOlmayanMessage()
	}
}

func FileAuthDelete(tokenString string, file models.File) (bool, map[string]interface{}) {
	tokenUser, status := repositories.GetCustomerByUserName(TokenParser(tokenString))
	var customer models.Customer
	customer = tokenUser.(models.Customer)
	if status {
		if customer.UserId == file.UserId {
			return true, service.OkMessage()
		} else {
			return false, service.YetkisiOlmayanMessage()
		}
	} else {
		return false, service.YetkisiOlmayanMessage()
	}
}
