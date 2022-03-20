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

func PriceById(priceId string) models.Price {
	var price models.Price
	config.DB.Find(&price, "id = ?", priceId)
	return price
}

func PriceAdd(price models.Price) {
	config.DB.Create(&price)
}

func PriceDelete(priceId string) {
	config.DB.Delete(models.Price{}, "id = ?", priceId)
}
func PriceUpdate(price models.Price) {
	config.DB.Save(&price)
}
