package auth

import (
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/daddydemir/kirtasiye-projesi/service"
)

func StationeryAuth(tokenString string, stationery models.Stationery) (bool, map[string]interface{}) {
	tokenStationery, status := repositories.GetStationeryByName(TokenParser(tokenString))
	var tempStationery models.Stationery
	tempStationery = tokenStationery.(models.Stationery)
	if status {
		if tempStationery.UserId == stationery.UserId {
			return true, service.OkMessage()
		} else {
			return false, service.YetkisiOlmayanMessage()
		}
	} else {
		return false, service.YetkisiOlmayanMessage()
	}
}
