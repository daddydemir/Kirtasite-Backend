package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func AllPrices() (interface{}, bool) {
	var prices []models.Price
	result := config.DB.Find(&prices)
	return returnModel(result.Error, prices)
}
func GetPriceById(priceId string) (interface{}, bool) {
	var price models.Price
	result := config.DB.Find(&price, "id = ?", priceId)
	return returnModel(result.Error, price)
}
func GetPriceByStationeryId(stationeryId string) (interface{}, bool) {
	var prices []models.Price
	result := config.DB.Find(&prices, "stationery_id = ?", stationeryId)
	return returnModel(result.Error, prices)
}
func AddPrice(price models.Price) (interface{}, bool) {
	result := config.DB.Create(&price)
	return returnModel(result.Error, "")
}
func DeletePrice(priceId string) (interface{}, bool) {
	result := config.DB.Delete(models.Price{}, "id = ?", priceId)
	return returnModel(result.Error, "")
}
func UpdatePrice(price models.Price) (interface{}, bool) {
	result := config.DB.Save(&price)
	return returnModel(result.Error, price)
}
