package auth

import (
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
)

func OrderAuthUser(tokenString string, userId int) (bool, map[string]string) {
	tokenUser, status := repositories.UserByName(TokenParser(tokenString))
	if status {
		if tokenUser.Id == userId {
			return true, map[string]string{"message": "Kullanıcı gereken yetkilere sahip."}
		} else {
			return false, map[string]string{"message": "Yetkisiz kullancı."}
		}
	} else {
		return false, map[string]string{"message": "Yetkisiz kullanıcı."}
	}
}

func OrderAuthStationery(tokenString string, stationeryId int) (bool, map[string]string) {
	tokenStationery, status := repositories.StationeryByName(TokenParser(tokenString))
	if status {
		if tokenStationery.Id == stationeryId {
			return true, map[string]string{"message": "Kullnıcı gereken yetkilere sahip."}
		} else {
			return false, map[string]string{"message": "Yetkisiz kullancı."}
		}
	} else {
		return false, map[string]string{"message": "Yetkisiz kullancı."}
	}
}

func OrderAuth(tokenString string, order models.Order) (bool, map[string]string) {
	tokenUser, status := repositories.UserByName(TokenParser(tokenString))
	if status {
		if tokenUser.Id == order.UserId {
			return true, map[string]string{"message": "Kullanıcı gereken yetkilere sahip."}
		} else {
			return false, map[string]string{"message": "Yetkisiz kullanıcı."}
		}
	} else {
		return false, map[string]string{"message": "Yetkisiz kullanıcı."}
	}
}
