package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func AllRoles() (interface{}, bool) {
	var roles []models.Role
	result := config.DB.Find(&roles)
	return returnModel(result.Error, roles)
}
