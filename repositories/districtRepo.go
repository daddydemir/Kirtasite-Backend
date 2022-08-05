package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func GetDistrictById(id string) (interface{}, bool) {
	var district models.District
	result := config.DB.Find(&district, "id = ?", id)
	return returnModel(result.Error, district)
}

func GetDistrictByCityId(id string) (interface{}, bool) {
	var district []models.District
	result := config.DB.Find(&district, "city_id = ?", id)
	return returnModel(result.Error, district)
}
