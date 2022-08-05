package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func GetAddressById(id string) (interface{}, bool) {
	var address models.Addresses
	result := config.DB.Find(&address, "id = ?", id)
	return returnModel(result.Error, address)
}
func GetAddressByCityId(id string) (interface{}, bool) {
	var address []models.Addresses
	result := config.DB.Find(&address, "city_id = ?", id)
	return returnModel(result.Error, address)
}
func GetAddressByDistrictId(id string) (interface{}, bool) {
	var address []models.Addresses
	result := config.DB.Find(&address, "district_id = ?", id)
	return returnModel(result.Error, address)
}
func AddAddress(address models.Address) (interface{}, bool) {
	result := config.DB.Create(&address)
	return returnModel(result.Error, address)
}
func UpdateAddress(address models.Addresses) (interface{}, bool) {
	result := config.DB.Save(&address)
	return returnModel(result.Error, address)
}
