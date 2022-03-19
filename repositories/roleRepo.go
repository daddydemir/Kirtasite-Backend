package repositories

import (
	"demir/config"
	"demir/models"
)

func RoleGetAll() []models.Role {

	var roles []models.Role
	config.DB.Find(&roles)
	return roles
}
