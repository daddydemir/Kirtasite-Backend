package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func RoleGetAll() []models.Role {

	var roles []models.Role
	config.DB.Find(&roles)
	return roles
}
