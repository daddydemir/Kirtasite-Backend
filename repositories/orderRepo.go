package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
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

func OrderById(id string) models.Order {
	var order models.Order
	config.DB.Find(&order, "id = ?", id)
	return order
}
