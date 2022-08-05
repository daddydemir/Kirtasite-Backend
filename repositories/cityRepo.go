package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
)

func GetCityById(id string) (interface{}, bool) {
	var city models.City
	result := config.DB.Find(&city, "id = ?", id)
	return returnModel(result.Error, city)
}
