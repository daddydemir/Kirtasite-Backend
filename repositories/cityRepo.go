package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func GetCityById(id string) models.City {
	var city models.City
	config.DB.Find(&city, "id = ?", id)
	return city
}
