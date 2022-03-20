package repositories

import (
	"demir/config"
	"demir/models"
)

func OrderByUserId(userId string) []models.Order {
	var orders []models.Order
	config.DB.Find(&orders, "user_id = ?", userId)
	return orders
}

func OrderByStationerId(stationeryId string) []models.Order {
	var orders []models.Order
	config.DB.Find(&orders, "stationery_id = ?", stationeryId)
	return orders
}

func OrderAdd(order models.Order) {
	config.DB.Create(&order)
}
