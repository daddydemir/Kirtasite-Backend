package auth

import (
	"demir/models"
	"demir/repositories"
)

func PriceAuth(tokenString string, price models.Price) (bool, map[string]string) {
	tokenStationery, status := repositories.StationeryByName(TokenParser(tokenString))
	if status {
		if price.StationeryId == tokenStationery.Id {
			return true, map[string]string{"message": "Kullanıcı gereken yetkilere sahip."}
		} else {
			return false, map[string]string{"message": "Yetkisiz kullanıcı."}
		}
	} else {
		return false, map[string]string{"message": "Yetkisiz kullanıcı."}
	}
}
