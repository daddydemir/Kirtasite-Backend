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

func StationeryGetById(stationeryId string) models.Stationery {
	var stationery models.Stationery
	config.DB.Find(&stationery, "id = ?", stationeryId)
	return stationery
}

func StationeryAdd(stationery models.Stationery) {
	config.DB.Create(&stationery)
}

func StationeryUpdate(stationery models.Stationery) {
	config.DB.Save(&stationery)
}

func StationeryDelete(stationeryId string) {
	config.DB.Delete(models.Stationery{}, "id = ?", stationeryId)
}
