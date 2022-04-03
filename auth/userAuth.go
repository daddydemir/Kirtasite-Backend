package auth

import (
	"demir/models"
	"demir/repositories"
)

func UserAuth(tokenString string, user models.User) (bool, map[string]string) {
	tokenUser, status := repositories.UserByName(TokenParser(tokenString))
	if status {
		if tokenUser.Id == user.Id {
			return true, map[string]string{"message": "Kullanıcı gereken yetkilere sahip."}
		} else {
			return false, map[string]string{"message": "Yetkisiz kullanıcı."}
		}
	} else {
		return false, map[string]string{"message": "Yetkisiz kullanıcı."}
	}
}
