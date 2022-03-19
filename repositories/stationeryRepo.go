package repositories

import (
	"demir/config"
	"demir/models"
)

func StationeryGetAll() []models.Stationery {
	var stationery []models.Stationery
	config.DB.Find(&stationery)
	return stationery
}
