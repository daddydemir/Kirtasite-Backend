package repositories

import (
	"demir/config"
	"demir/models"
)

func GetOrderByUserId() []models.Order {
	var orders []models.Order
	config.DB.Find(&orders, "user_id = ?", 9)
	return orders
}
