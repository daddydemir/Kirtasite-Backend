package repositories

import (
	"demir/config"
	"demir/models"
)

func GetAllPrices() []models.Price {
	var prices []models.Price
	config.DB.Find(&prices)
	return prices
}
