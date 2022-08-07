package auth

import (
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/daddydemir/kirtasiye-projesi/service"
)

func PriceAuth(tokenString string, price models.Price) (bool, map[string]interface{}) {
	tokenStationery, status := repositories.GetStationeryByName(TokenParser(tokenString))
	var tempStationery models.Stationery
	tempStationery = tokenStationery.(models.Stationery)
	if status {
		if price.StationeryId == tempStationery.UserId {
			return true, service.OkMessage()
		} else {
			return false, service.YetkisiOlmayanMessage()
		}
	} else {
		return false, service.YetkisiOlmayanMessage()
	}
}
