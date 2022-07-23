package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func GetDistrictById(id string) models.District {
	var district models.District
	config.DB.Find(&district, "id = ?", id)
	return district
}

func GetDistrictByCity(id string) []models.District {
	var district []models.District
	config.DB.Find(&district, "city_id = ?", id)
	return district
}
