package auth

import (
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
)

func FileAuthUser(tokenString string, userId int) (bool, map[string]string) {
	tokenUser, status := repositories.UserByName(TokenParser(tokenString))
	if status {
		if tokenUser.Id == userId {
			return true, map[string]string{"message": "Kullanıcı gereken yetkilere sahip."}
		} else {
			return false, map[string]string{"message": "Yetkisiz kullanıcı."}
		}
	} else {
		return false, map[string]string{"message": "Yetkisiz kullanıcı."}
	}
}

func FileAuthDelete(tokenString string, file models.File) (bool, map[string]string) {
	tokenUser, status := repositories.UserByName(TokenParser(tokenString))
	if status {
		if tokenUser.Id == file.UserId {
			return true, map[string]string{"message": "Kullanıcı gereken yetkilere sahip."}
		} else {
			return false, map[string]string{"message": "Yetkisiz kullanıcı."}
		}
	} else {
		return false, map[string]string{"message": "Yetksisiz kullanıcı."}
	}
}
