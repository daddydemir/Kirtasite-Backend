package auth

import (
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/daddydemir/kirtasiye-projesi/service"
)

func OrderAuthUser(tokenString string, userId int) (bool, map[string]interface{}) {
	tokenUser, status := repositories.GetCustomerByUserName(TokenParser(tokenString))
	var tempUser models.User
	tempUser = tokenUser.(models.User)
	if status {
		if tempUser.Id == userId {
			return true, service.OkMessage()
		} else {
			return false, service.YetkisiOlmayanMessage()
		}
	} else {
		return false, service.YetkisiOlmayanMessage()
	}
}

func OrderAuthStationery(tokenString string, stationeryId int) (bool, map[string]interface{}) {
	tokenStationery, status := repositories.GetStationeryByName(TokenParser(tokenString))
	var tempStationery models.Stationery
	tempStationery = tokenStationery.(models.Stationery)
	if status {
		if tempStationery.UserId == stationeryId {
			return true, service.OkMessage()
		} else {
			return false, service.YetkisiOlmayanMessage()
		}
	} else {
		return false, service.YetkisiOlmayanMessage()
	}
}

func OrderAuth(tokenString string, order models.Order) (bool, map[string]interface{}) {
	tokenUser, status := repositories.GetCustomerByUserName(TokenParser(tokenString))
	var customer models.Customer
	customer = tokenUser.(models.Customer)
	if status {
		if customer.UserId == order.CustomerId {
			return true, service.OkMessage()
		} else {
			return false, service.YetkisiOlmayanMessage()
		}
	} else {
		return false, service.YetkisiOlmayanMessage()
	}
}
