package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func GetOrderByCustomerId(userId string) (interface{}, bool) {
	var orders []models.Orders
	result := config.DB.Find(&orders, "customer_id = ?", userId)
	return returnModel(result.Error, orders)
}
func GetOrderByStationeryId(stationeryId string) (interface{}, bool) {
	var orders []models.Orders
	result := config.DB.Find(&orders, "stationery_id = ?", stationeryId)
	return returnModel(result.Error, orders)
}
func AddOrder(order models.Order) (interface{}, bool) {
	result := config.DB.Create(&order)
	return returnModel(result.Error, order)
}
func GetOrderById(id string) (interface{}, bool) {
	var order models.Orders
	result := config.DB.Find(&order, "id = ?", id)
	return returnModel(result.Error, order)
}
func UpdateOrder(order models.Order) (interface{}, bool) {
	result := config.DB.Save(&order)
	return returnModel(result.Error, order)
}
