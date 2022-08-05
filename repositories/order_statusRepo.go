package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func GetOrderStatusById(id string) (interface{}, bool) {
	var orderStatus models.Order_Status
	result := config.DB.Find(&orderStatus, "id = ?", id)
	return returnModel(result.Error, orderStatus)
}
