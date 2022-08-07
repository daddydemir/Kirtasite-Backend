package auth

import (
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/daddydemir/kirtasiye-projesi/service"
)

func UserAuth(tokenString string, user models.Users) (bool, map[string]interface{}) {
	tokenUser, status := repositories.GetCustomerByUserName(TokenParser(tokenString))
	var tempUser models.Users
	tempUser = tokenUser.(models.Users)
	if status {
		if tempUser.Id == user.Id {
			return true, service.OkMessage()
		} else {
			return false, service.YetkisiOlmayanMessage()
		}
	} else {
		return false, service.YetkisiOlmayanMessage()
	}
}
