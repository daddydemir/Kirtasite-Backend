package auth

import (
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/daddydemir/kirtasiye-projesi/service"
)

func CommentAuth(tokenString string, userId int) (bool, map[string]interface{}) {
	status, message := IsValid(tokenString)
	if status {
		tokenUser, state := repositories.GetCustomerByUserId(TokenParser(tokenString))
		var customer models.Customer
		customer = tokenUser.(models.Customer)
		if state {
			if customer.UserId == userId {
				return true, service.OkMessage()
			} else {
				return false, service.YetkisiOlmayanMessage()
			}
		} else {
			return false, service.YetkisiOlmayanMessage()
		}
	} else {
		return false, message
	}
}
