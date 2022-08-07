package auth

import (
	"fmt"

	"github.com/daddydemir/kirtasiye-projesi/repositories"
	"github.com/daddydemir/kirtasiye-projesi/security"
)

func LoginToSystem(username string, password string) (map[string]string, bool) {
	user, status := repositories.UserByName(username)
	stationery, state := repositories.StationeryByName(username)
	if !status && !state {
		return map[string]string{"message": "Kullanıcı adı yada Parola hatalı."}, false
	} else if status {
		isSuccess := security.CheckPassword(user.Password, password)
		if isSuccess {
			token := GenerateToken(user.Username)
			return map[string]string{
				"id":       fmt.Sprint(user.Id),
				"username": user.Username,
				"mail":     user.Mail,
				"phone":    user.TelNo,
				"image":    user.ImagePath,
				"role":     "STUDENT",
				"token":    token}, true
		} else {
			return map[string]string{"message": "Kullanıcı adı yada Parola hatalı."}, false
		}
	} else if state {
		isSuccess := security.CheckPassword(stationery.Password, password)
		if isSuccess {
			token := GenerateToken(stationery.CompanyName)
			return map[string]string{
				"id":       fmt.Sprint(stationery.Id),
				"username": stationery.CompanyName,
				"mail":     stationery.Mail,
				"phone":    stationery.TelNo,
				"image":    stationery.ImagePath,
				"role":     "STATIONERY",
				"token":    token}, true
		} else {
			return map[string]string{"message": "Kullanıcı adı yada Parola hatalı."}, false
		}
	} else {
		return map[string]string{"message": "Beklenmedik bir hata oluştu, daha sonra tekrar deneyiniz"}, false
	}
}
